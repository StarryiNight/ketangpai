package controller

import (
	"fmt"
	"ketangpai/dao/redis"
	"ketangpai/logic"
	"ketangpai/models"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CreatePostHandler  发表帖子
func CreatePostHandler(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	// 获取作者ID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	post.AuthorId = userID

	err = logic.CreatePost(&post)
	if err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// PostListHandler 帖子列表(获取投票信息等）
func PostListHandler(c *gin.Context) {
	order, _ := c.GetQuery("order")
	pageStr, ok := c.GetQuery("page")
	if !ok {
		pageStr = "1"
	}
	pageNum, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		pageNum = 1
	}
	posts := redis.GetPost(order, pageNum)
	fmt.Println(len(posts))
	ResponseSuccess(c, posts)
}

// PostList2Handler （帖子状态等详细信息）
func PostList2Handler(c *gin.Context) {
	data, err := logic.GetPostList2()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)

}

// PostDetailHandler 帖子详情
func PostDetailHandler(c *gin.Context) {
	postId := c.Param("id")

	post, err := logic.GetPost(postId)
	if err != nil {
		zap.L().Error("logic.GetPost(postID) failed", zap.String("postId", postId), zap.Error(err))
	}

	ResponseSuccess(c, post)
}
