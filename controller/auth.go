package controller

import (
	"errors"
	"fmt"
	"ketangpai/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ContextUserIDKey    = "userID"
	ContextUserPosition = "position"
)

var (
	ErrorUserNotLogin = errors.New("当前用户未登录")
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Token在Header的Authorization中，Bearer开头
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
		// parts[1]获取到tokenString，用解析JWT的函数来解析
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			fmt.Println(err)
			ResponseError(c, CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set(ContextUserIDKey, mc.UserId)
		c.Set(ContextUserPosition,mc.Position)
		c.Next()
	}
}
