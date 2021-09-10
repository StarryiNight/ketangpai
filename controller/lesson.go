package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ketangpai/dao/mysql"
	"ketangpai/dao/redis"
	"ketangpai/logic"
	"ketangpai/models"
	"strconv"
)

// CreateLessonHandler   创建课堂
func CreateLessonHandler(c *gin.Context) {
	var lesson models.Lesson
	lesson.TeacherId, _ = GetCurrentUserID(c)
	var err error
	lesson.ClassId,err=strconv.ParseInt(c.Param("classid"),10,64)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,CodeInvalidParams.Msg())
	}
	// 创建帖子
	if err = mysql.CreateLesson(&lesson); err != nil {
		zap.L().Error("mysql.CreateLesson(&lesson) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)

}

// SetHomeworkHandler 布置作业
func SetHomeworkHandler(c *gin.Context) {
	homework:=c.PostForm("homework")
	lessonid,err:=strconv.ParseInt(c.Param("lessonid"),10,64)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,CodeInvalidParams.Msg())
	}

	if err = mysql.SetHomework(homework,lessonid); err != nil {
		zap.L().Error("mysql.SetHomework(homework,lessonid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// LessonOverHandler 下课
func LessonOverHandler(c *gin.Context) {
	lessonid,err:=strconv.ParseInt(c.Param("lessonid"),10,64)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,CodeInvalidParams.Msg())
	}
	if err = mysql.LessonOver(lessonid); err != nil {
		zap.L().Error("mysql.LessonOver(lessonid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)

}

// LessonSignInHandler 学生签到
func LessonSignInHandler(c *gin.Context) {
	var lessonInfo models.LessonInfo
	var err error
	lessonInfo.LessonId,err=strconv.ParseInt(c.Param("lessonid"),10,64)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,CodeInvalidParams.Msg())
	}
	lessonInfo.StudentId, _ = GetCurrentUserID(c)

	err = logic.LessonSignIn(&lessonInfo)
	if err != nil {
		zap.L().Error("logic.LessonSignIn failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy,err.Error())
		return
	}
	ResponseSuccess(c, nil)
}

// HomeworkSubmitHandler 作业提交
func HomeworkSubmitHandler(c *gin.Context) {
	lessonid,err:=strconv.ParseInt(c.Param("lessonid"),10,64)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,CodeInvalidParams.Msg())
	}
	studentid,_:= GetCurrentUserID(c)
	if err := mysql.HomeworkSubmit(lessonid,studentid); err != nil {
		zap.L().Error("mysql.HomeworkSubmit(lessonid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// LessonCheckHandler 查看课堂作业完成情况
func LessonCheckHandler(c *gin.Context) {
	lessonid,err:=strconv.ParseInt(c.Param("lessonid"),10,64)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,CodeInvalidParams.Msg())
	}
	lessoninfos,err := mysql.LessonCheck(lessonid)
	if  err != nil {
		zap.L().Error("mysql.LessonCheck(lessonid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, lessoninfos)
}

// LessonTalkRankHandler 获取课堂发言排行榜（根据页数）
func LessonTalkRankHandler(c *gin.Context) {
	lessonid:=c.Param("lessonid")
	page, err:=strconv.ParseInt(c.Param("page"),10,64)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,CodeInvalidParams.Msg())
	}

	rank, err := redis.GetTalkRank(page, lessonid)
	if err != nil {
		zap.L().Error("redis.GetTalkRank(page, lessonid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c,rank)
	//ResponseSuccess(c,lessonid)
	//ResponseSuccess(c,page)
}