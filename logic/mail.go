package logic

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"ketangpai/dao/redis"
	"math/rand"
	"strconv"
	"time"
)

const (
	user = "******"
	pass="*******"
	host="smtp.qq.com"

)

func SendMail(username string,address string)error {

	//生成6位随机验证码
	randPass:= rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000000)

	body:=fmt.Sprintf( "尊敬的用户:%s 您正在使用课堂派邮箱找回密码功能  验证码是：%d。该验证码有效时间为30分钟，如不是您本人操作，请登录课堂派管理账户安全！",username,randPass)
	subject := "课堂派找回密码"

	//储存验证码到redis中
	err := redis.SetVerificationCode(username, strconv.Itoa(randPass))
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From","ketangpai"+"<"+user+">")
	m.SetHeader("To", address)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)


	d := gomail.NewDialer(host,465, user, pass)
	err = d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}
