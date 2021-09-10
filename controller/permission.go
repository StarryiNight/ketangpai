package controller

import (
	"github.com/gin-gonic/gin"
	"ketangpai/dao/mysql"
)

func PermissionMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取claim
		position, err:= GetCurrentUserPosition(c)

		// 检查用户权限
		isPass, err := mysql.Enforcer.Enforce(position, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			ResponseErrorWithMsg(c, CodeLimitedAuthority, CodeLimitedAuthority.Msg())
			c.Abort()
			return
		}
		if isPass {
			c.Next()
		} else {
			ResponseErrorWithMsg(c, CodeLimitedAuthority, CodeLimitedAuthority.Msg())
			c.Abort()
			return
		}
	}
}