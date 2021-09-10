package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"ketangpai/dao/redis"

	"github.com/gin-gonic/gin"
)

type VoteData struct {
	PostID    string  `json:"post_id"`
	//1赞成 -1反对 0未投票
	Direction float64 `json:"direction"`
}

func (v *VoteData) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		PostID    string  `json:"post_id"`
		Direction float64 `json:"direction"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.PostID) == 0 {
		err = errors.New("缺少必填字段post_id")
	} else if required.Direction == 0 {
		err = errors.New("缺少必填字段direction")
	} else {
		v.PostID = required.PostID
		v.Direction = required.Direction
	}
	return
}

// VoteHandler 为帖子投票
func VoteHandler(c *gin.Context) {
	var vote VoteData
	if err := c.ShouldBindJSON(&vote); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	//获取当前登陆的id
	userID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNotLogin)
		return
	}
	//把用户为帖子的投票记录保存到redis中

	if err := redis.PostVote(vote.PostID, fmt.Sprint(userID), vote.Direction); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "投票成功")
}
