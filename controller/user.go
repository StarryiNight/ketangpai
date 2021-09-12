package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"ketangpai/dao/mysql"
	"ketangpai/dao/redis"
	"ketangpai/logic"
	"ketangpai/models"
	"ketangpai/pkg/jwt"
	"net/http"
	"strings"
)

// SignUpHandler 注册
func SignUpHandler(c *gin.Context) {
	//序列化 绑定到定义的注册请求参数结构体上
	var fo models.RegisterForm
	if err := c.ShouldBindJSON(&fo); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		c.JSON(http.StatusOK, gin.H{
			"msg": RemoveTopStruct(errs.Translate(trans)),
		})
		return

	}
	// 注册用户 插入到数据库中
	err := mysql.Register(&models.User{
		UserName: fo.UserName,
		Password: fo.Password,
		Email:    fo.Email,
	})
	if errors.Is(err, mysql.ErrorUserExit) {
		ResponseError(c, CodeUserExist)
		return
	}
	if err != nil {
		zap.L().Error("mysql.Register() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	//序列化 绑定到定义的登陆请求参数结构体上
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}

	//查询数据库 比对账户密码 返回结果
	if err := mysql.Login(&u); err != nil {
		zap.L().Error("mysql.Login(&u) failed", zap.Error(err))
		ResponseError(c, CodeInvalidPassword)
		return
	}

	// 生成Token
	//返回一个access登陆token，和一个本地用来刷新的token
	aToken, rToken, _ := jwt.GenToken(u)

	ResponseSuccess(c, gin.H{
		"4.accessToken":  aToken,
		"5.refreshToken": rToken,
		"3.position":     u.Position,
		"1.userID":       u.UserID,
		"2.username":     u.UserName,
	})
}

// RefreshTokenHandler 用本地上传的refreshToken刷新token重新生成携带登陆信息的accesstoken
func RefreshTokenHandler(c *gin.Context) {
	rt := c.Query("refresh_token")
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		ResponseErrorWithMsg(c, CodeInvalidToken, "请求头缺少Auth Token")
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseErrorWithMsg(c, CodeInvalidToken, "Token格式不对")
		c.Abort()
		return
	}
	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}

// RetrievePasswordHandler1 发送找回密码的验证码
func RetrievePasswordHandler1(c *gin.Context) {
	username := c.PostForm("username")
	err := logic.RetrievePassword(username)
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(c, "发送验证码成功")
}

func RetrievePasswordHandler2(c *gin.Context) {
	username := c.Param("username")
	var p models.NewPassword
	if err := c.ShouldBind(&p); err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	//获取存储在redis的验证码
	code, err := redis.GetVerificationCode(username)
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, CodeServerBusy.Msg())
		return
	}
	if code != p.Code {
		ResponseErrorWithMsg(c, CodeVerificationCode, CodeVerificationCode.Msg())
		return
	}
	if err := mysql.ChangePassword(username, p.Password); err != nil {
		zap.L().Error("mysql.ChangePassword(username,p.Password);", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, CodeServerBusy.Msg())
		return
	}
	err = redis.DelVerificationCode(username)
	if err != nil {
		return
	}
	ResponseSuccess(c, "修改成功")
}
