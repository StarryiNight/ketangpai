package routers

import (
	"ketangpai/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/chat/:roomid/:channel/:token", controller.Server)
	v1 := r.Group("/api/v1")
	v1.POST("/login", controller.LoginHandler)
	v1.POST("/signup", controller.SignUpHandler)
	//v1.GET("/refresh_token", controller.RefreshTokenHandler)


	v1.Use(controller.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.PostDetailHandler)
		v1.GET("/posts", controller.PostListHandler)

		v1.GET("/posts2", controller.PostList2Handler)

		v1.POST("/vote", controller.VoteHandler)

		v1.POST("/comment", controller.CommentHandler)
		v1.GET("/comment", controller.CommentListHandler)



		//v1.GET("/ping", func(c *gin.Context) {
		//	c.String(http.StatusOK, "pong")
		//})

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
