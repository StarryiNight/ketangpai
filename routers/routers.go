package routers

import (
	"ketangpai/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//上课讨论、弹幕、抽问、抢答
	r.GET("/chat/:lessonid/:channel/:token", controller.Server)
	v1 := r.Group("/api/v1")
	//登陆注册
	v1.POST("/login", controller.LoginHandler)
	v1.POST("/signup", controller.SignUpHandler)

	//用refresh token刷新access token
	v1.GET("/refresh_token", controller.RefreshTokenHandler)



	v1.Use(controller.JWTAuthMiddleware())
	{
		//上传业务路由
		v1.POST("/upload",controller.UploadHandler)
		//下载业务路由
		v1.GET("/download",controller.DownloadHandler)


		//添加课程
		v1.POST("/class",controller.ClassAddHandler)
		//增加成绩
		v1.POST("/score/:classid",controller.ScoreAddHandler)
		//学生选课
		v1.POST("/studentAddClass/:classid",controller.StudentAddClassHandler)

		//创建课堂
		v1.POST("/lesson/create/:classid",controller.CreateLessonHandler)
		//布置作业
		v1.PATCH("/lesson/homework/:lessonid",controller.SetHomeworkHandler)
		//下课
		v1.PATCH("/lesson/over/:lessonid",controller.LessonOverHandler)
		//课堂情况查看
		v1.GET("/lesson/check/:lessonid",controller.LessonCheckHandler)
		//获取课堂发言情况
		v1.GET("/lesson/rank/:lessonid/:page",controller.LessonTalkRankHandler)


		//学生签到
		v1.POST("student/signin/:lessonid",controller.LessonSignInHandler)
		//学生提交作业
		v1.POST("student/homework/:lessonid",controller.HomeworkSubmitHandler)



		//获取板块列表
		v1.GET("/community", controller.CommunityHandler)
		//获取某个板块详情
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		//发表帖子
		v1.POST("/post", controller.CreatePostHandler)
		//获取某个帖子详情
		v1.GET("/post/:id", controller.PostDetailHandler)
		//获取帖子列表（投票等信息） 按热度或时间排序
		v1.GET("/posts", controller.PostListHandler)
		//获取帖子列表（状态等信息） 按热度或时间排序
		v1.GET("/posts2", controller.PostList2Handler)

		//为帖子投票
		v1.POST("/vote", controller.VoteHandler)

		//发表评论
		v1.POST("/comment", controller.CommentHandler)
		//获取帖子的评论
		v1.GET("/comment", controller.CommentListHandler)


		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
