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


	v1.Use(controller.JWTAuthMiddleware()).Use(controller.PermissionMiddleWare())
	//v1.Use(controller.JWTAuthMiddleware())
	{

		//修改身份权限
		v1.GET("/permission",controller.GetAllPolicyHandler)
		v1.POST("/permission",controller.AddPolicyHandler)

		//上传业务路由
		v1.POST("/upload",controller.UploadHandler)
		//下载业务路由
		v1.GET("/download",controller.DownloadHandler)

		
		class:=v1.Group("/class")
		{
			//添加课程
			class.POST("",controller.ClassAddHandler)
			//增加成绩
			class.POST("/score/:classid",controller.ScoreAddHandler)

		}

		lesson:=v1.Group("/lesson")
		{
			//创建课堂
			lesson.POST("/create/:classid",controller.CreateLessonHandler)
			//布置作业
			lesson.PATCH("/homework/:lessonid",controller.SetHomeworkHandler)
			//下课
			lesson.PATCH("/over/:lessonid",controller.LessonOverHandler)
			//课堂情况查看
			lesson.GET("/check/:lessonid",controller.LessonCheckHandler)
			//获取课堂发言情况
			lesson.GET("/rank/:lessonid/:page",controller.LessonTalkRankHandler)

		}
		
		student:=v1.Group("/student")
		{
			//学生签到
			student.POST("/signin/:lessonid",controller.LessonSignInHandler)
			//学生提交作业
			student.POST("/homework/:lessonid",controller.HomeworkSubmitHandler)
			//学生选课
			student.POST("/choseclass/:classid",controller.StudentAddClassHandler)
		}
		

		community:=v1.Group("community")
		{
			//获取板块列表
			community.GET("/forum", controller.CommunityHandler)
			//获取某个板块详情
			community.GET("/forum/:id", controller.CommunityDetailHandler)

			//发表帖子
			community.POST("/post", controller.CreatePostHandler)
			//获取某个帖子详情
			community.GET("/post/:id", controller.PostDetailHandler)
			//获取帖子列表（投票等信息） 按热度或时间排序
			community.GET("/posts", controller.PostListHandler)
			//获取帖子列表（状态等信息） 按热度或时间排序
			community.GET("/posts2", controller.PostList2Handler)

			//为帖子投票
			community.POST("/vote", controller.VoteHandler)

			//发表评论
			community.POST("/comment", controller.CommentHandler)
			//获取帖子的评论
			community.GET("/comment", controller.CommentListHandler)

		}
		

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
