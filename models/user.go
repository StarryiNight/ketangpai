package models

import "time"

// User 登陆api请求param结构
type User struct {
	UserID   int64  `json:"user_id" db:"user_id"`
	UserName string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Position string `json:"position" db:"position"`
	Email    string `json:"email" db:"email"`
}

// RegisterForm 注册api请求param结构
type RegisterForm struct {
	UserName   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required,eqfield=RePassword"`
	RePassword string `json:"repassword" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
}

type NewPassword struct {
	Password   string `json:"password" binding:"required,eqfield=RePassword"`
	RePassword string `json:"repassword" binding:"required"`
	Code       string `json:"code"`
}

type StudentScore struct {
	StudentName string    `json:"student_name" db:"student_name"`
	StudentID   int64     `json:"student_id" db:"student_id"`
	Score       float32   `json:"score" db:"score"`
	SubmitTime  time.Time `json:"submit_time" db:"submit_time"`
}
