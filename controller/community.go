package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ketangpai/dao/mysql"
)

// 板块

// CommunityHandler 板块列表
func CommunityHandler(c *gin.Context) {
	communityList, err := mysql.GetCommunityList()
	if err != nil {
		zap.L().Error("mysql.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, communityList)
}

// CommunityDetailHandler 板块详情
func CommunityDetailHandler(c *gin.Context) {
	communityID := c.Param("id")
	communityList, err := mysql.GetCommunityByID(communityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID() failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeSuccess, err.Error())
		return
	}
	ResponseSuccess(c, communityList)
}
