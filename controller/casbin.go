package controller

import (
	"github.com/gin-gonic/gin"
	"ketangpai/dao/mysql"
	"ketangpai/models"
)

func GetAllPolicyHandler(c *gin.Context) {
	policy:=mysql.Enforcer.GetPolicy()
	ResponseSuccess(c,policy)
}

func AddPolicyHandler(c *gin.Context) {
	var m models.Permission
	err := c.ShouldBind(&m)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,CodeInvalidParams.Msg())
		return
	}
	result, err := mysql.Enforcer.AddPolicy(m.Role, m.Path, m.Method)
	if err != nil ||!result{
		ResponseErrorWithMsg(c,CodeServerBusy,CodeServerBusy.Msg())
		return
	}
	ResponseSuccess(c,"success")
}
