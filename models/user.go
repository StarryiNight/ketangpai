package models

// User 登陆api请求param结构
type User struct {
	UserID   uint64 `json:"user_id" db:"user_id"`
	UserName string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

// RegisterForm 注册api请求param结构
type RegisterForm struct {
	UserName        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"repassword"`
}

