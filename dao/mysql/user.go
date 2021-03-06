package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"ketangpai/models"
	"ketangpai/pkg/snowflake"
)

const secret = "ketangpai"

//encryptPassword 给用户存储在数据库的密码加密
func encryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(data))
}

func Register(user *models.User) (err error) {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int64
	err = db.Get(&count, sqlStr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if count > 0 {
		// 用户已存在
		return ErrorUserExit
	}
	// 生成user_id
	userID, err := snowflake.GetID()
	if err != nil {
		return ErrorGenIDFailed
	}
	// 生成加密密码
	password := encryptPassword([]byte(user.Password))
	// 把用户插入数据库
	sqlStr = "insert into user(user_id, username, password,email) values (?,?,?,?)"
	_, err = db.Exec(sqlStr, userID, user.UserName, password,user.Email)
	return
}

func Login(user *models.User) (err error) {
	originPassword := user.Password // 记录原始密码
	sqlStr := "select user_id, username, password ,position from user where username = ?"
	err = db.Get(user, sqlStr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		// 查询数据库出错
		return
	}
	if err == sql.ErrNoRows {
		// 用户不存在
		return ErrorUserNotExit
	}
	// 生成加密密码与查询到的密码比较
	password := encryptPassword([]byte(originPassword))
	if user.Password != password {
		return ErrorPasswordWrong
	}
	return
}

func GetUserByID(idStr string) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username,position from user where user_id = ?`
	err = db.Get(user, sqlStr, idStr)
	return
}

func GetUserByName(user *models.User) (flag bool, err error) {
	sqlStr := `select user_id, username,position ,email from user where username = ?`
	err = db.Get(user, sqlStr, user.UserName)
	if err != nil {
		return false,err
	}
	if user.UserName!="" {
		return true,nil
	}else{
		return false,nil
	}
}

func ChangePassword(username string, password string) error {
	sqlStr := `update user set password=? where username=?`
	pswd := encryptPassword([]byte(password))
	_,err := db.Exec(sqlStr,pswd, username )
	if err != nil {
		return err
	}
	return nil
}