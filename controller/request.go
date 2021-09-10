package controller

import (
	"github.com/gin-gonic/gin"
)

// GetCurrentUserID 获取当前登陆的用户id
func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	_userID, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = _userID.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// GetCurrentUserPosition 获取当前登陆的用户身份
func GetCurrentUserPosition(c *gin.Context) (position string, err error) {
	_position, ok := c.Get(ContextUserPosition)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	position, ok = _position.(string)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}