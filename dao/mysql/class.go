package mysql

import (
	"go.uber.org/zap"
	"ketangpai/models"
	"log"
)

// CreateClass  发表评论
func CreateClass(class *models.Class) (err error) {
	sqlStr := `insert into class (class_id, class_name)
	values(?,?)`
	_, err = db.Exec(sqlStr, class.ClassID,class.ClassName)
	if err != nil {
		zap.L().Error("insert class failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

func AddScore(classid int64,userid int64,score int) (err error) {
	sqlStr := `UPDATE score SET score=(?) WHERE class_id=(?) and user_id=?`
	_, err = db.Exec(sqlStr, score,classid,userid)
	if err != nil {
		zap.L().Error("Add Score failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

func StudentAddClass(classid int64, userid int64) (err error ){
	sqlStr := `insert into score (class_id, user_id)
	values(?,?)`
	log.Println(classid,userid)
	_, err = db.Exec(sqlStr, classid,userid)
	if err != nil {
		zap.L().Error("studentAddClass failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}
