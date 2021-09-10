package logic

import (
	"errors"
	"ketangpai/dao/mysql"
	"ketangpai/models"
)

func RetrievePassword(username string) (err error) {
	user := new(models.User)
	user.UserName=username
	flag, err :=mysql.GetUserByName(user)
	if err != nil {
		return err
	}
	if !flag {
		return errors.New("没有此用户")
	}
	err=SendMail(user.UserName,user.Email)
	if err != nil {
		return err
	}
	return nil
}
