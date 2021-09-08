package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ketangpai/dao/mysql"
	"ketangpai/models"
	"strconv"
)

// ClassAddHandler  添加课程
func ClassAddHandler(c *gin.Context) {
	var class models.Class
	if err := c.BindJSON(&class); err != nil {
		fmt.Println(err)
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 创建帖子
	if err := mysql.CreateClass(&class); err != nil {
		zap.L().Error("mysql.CreateClass(&class) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// ScoreAddHandler 添加学生分数
func ScoreAddHandler(c *gin.Context) {
	classid, _ :=strconv.ParseInt(c.Param("classid"),10,64)
	userid, _ :=strconv.ParseInt(c.PostForm("user_id"),10,64)
	score, _ :=strconv.Atoi(c.PostForm("score"))
	err := mysql.AddScore(classid, int64(userid),score)
	if err != nil {
		zap.L().Error("mysql.AddScore failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// StudentAddClassHandler  学生选课
func StudentAddClassHandler(c *gin.Context) {
	userid, _ :=getCurrentUserID(c)
	classid, _ :=strconv.ParseInt(c.Param("classid"),10,64)

	err := mysql.StudentAddClass(classid,userid)
	if err != nil {
		zap.L().Error("mysql.StudentAddClass failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}