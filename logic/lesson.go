package logic

import (
	"go.uber.org/zap"
	"ketangpai/dao/mysql"
	"ketangpai/models"
	"time"
)

func LessonSignIn(lessonInfo *models.LessonInfo) (err error) {
	lesson, err := mysql.GetLessonById(lessonInfo.LessonId)
	if err != nil {
		zap.L().Error("mysql.GetLessonById(lessonInfo.LessonId) failed", zap.Int64("lessonId", lessonInfo.LessonId), zap.Error(err))
		return  err
	}
	lessonInfo.SignInTime=time.Now()
	lessonInfo.SignInStatus=1
	//u, _ := time.ParseInLocation("2006-01-02 15:04:05", "2030-01-01 00:00:00",time.Local)
	if !lesson.EndTime.IsZero(){
		zap.L().Error(" 1", zap.Time("a",lesson.EndTime))
		if lessonInfo.SignInTime.After(lesson.EndTime){
			zap.L().Error(" 2", zap.Time("a",lesson.EndTime))
			lessonInfo.SignInStatus=-1
		}
	}
	err = mysql.LessonSignIn(lessonInfo)
	if err != nil {
		zap.L().Error(" mysql.LessonSignIn(lessonInfo) failed", zap.Error(err))
		return err
	}
	return
}

