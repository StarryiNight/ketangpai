package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"ketangpai/models"
	"ketangpai/pkg/snowflake"
	"time"
)

// CreateLesson 创建课堂
func CreateLesson(lesson *models.Lesson) (err error) {
	lesson.LessonId, _ = snowflake.GetID()
	sqlStr := `insert into lesson(lesson_id, class_id, teacher_id)
	values(?,?,?)`
	_, err = db.Exec(sqlStr, lesson.LessonId, lesson.ClassId, lesson.TeacherId)
	if err != nil {
		zap.L().Error("insert lesson failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

func SetHomework(homework string, lessonid int64) (err error) {
	sqlStr := `update lesson set homework=? where lesson_id=?`
	_, err = db.Exec(sqlStr, homework, lessonid)
	if err != nil {
		zap.L().Error("update lesson homework failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

// LessonOver 下课
func LessonOver(lessonid int64) (err error) {
	sqlStr := `update lesson set end_time=? where lesson_id=?`
	_, err = db.Exec(sqlStr, time.Now(), lessonid)
	if err != nil {
		zap.L().Error("update lesson end_time failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

// LessonSignIn  签到
func LessonSignIn(lessonInfo *models.LessonInfo) (err error) {
	sqlStr := `Insert into lessoninfo( lesson_id, student_id, signin_status, signin_time)values(?,?,?,?)`
	_, err = db.Exec(sqlStr, lessonInfo.LessonId, lessonInfo.StudentId, lessonInfo.SignInStatus, lessonInfo.SignInTime)
	if err != nil {
		zap.L().Error("insert lessoninfo failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

// GetLessonById 通过lessonId获取课堂情况
func GetLessonById(lessonid int64) (lesson *models.Lesson, err error) {
	sql1 := `SELECT end_time is null  FROM lesson where lesson_id=?`
	var flag int
	db.Get(&flag, sql1, lessonid)

	lesson = new(models.Lesson)
	var sqlStr string
	if flag == 1 {
		sqlStr=`select lesson_id,class_id,teacher_id,ifnull(homework,'空') as homework,start_time 
			from lesson 
			where lesson_id=?`
	} else {
		sqlStr=`select lesson_id,class_id,teacher_id,ifnull(homework,'空') as homework,start_time,end_time
			from lesson 
			where lesson_id=?`
	}
	err = db.Get(lesson, sqlStr, lessonid)

	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query lesson failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	return
}

// HomeworkSubmit 提交作业 修改状态和时间
func HomeworkSubmit(lessonid int64,studentid int64) (err error) {
	sqlStr := `update  lessoninfo set submit_status=? , submit_time=? where student_id=? and lesson_id=?`
	_, err = db.Exec(sqlStr, 1,time.Now(),studentid,lessonid)
	if err != nil {
		zap.L().Error("update lessoninfo homework failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}



// LessonCheck 根据lessonid查看课堂签到、作业完成情况
func LessonCheck(lessonid int64) (lessoninfos []*models.LessonInfo, err error) {
	sqlStr:= `SELECT a.lesson_id,a.student_id,username as student_name,a.submit_status,submit_time,signin_status,signin_time
				FROM(SELECT lesson_id,student_id,submit_status,submit_time,signin_status,signin_time
						FROM lessoninfo
						where lesson_id=?)AS a
				LEFT JOIN user
				ON user_id=a.student_id`
	err=db.Select(&lessoninfos,sqlStr, lessonid)
	if err != nil {
		zap.L().Error("query lessoninfo  failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	return
}

func CheckSignIn(lessonInfo *models.LessonInfo) (bool,error) {
	sqlStr:= `select signin_status from lessoninfo where lesson_id=? and student_id=?`
	var flag int
	err:=db.Get(&flag,sqlStr,lessonInfo.LessonId,lessonInfo.StudentId)
	if err != nil {
		zap.L().Error("query lessoninfo  signin_status failed", zap.Error(err))
		return false,err
	}
	if flag!= 0 {
		return true,nil
	}else {
		return false,nil
	}
}