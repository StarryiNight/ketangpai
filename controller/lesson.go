package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"ketangpai/dao/mysql"
	"ketangpai/dao/redis"
	"ketangpai/logic"
	"ketangpai/models"
	"net/http"
	"strconv"
)

// CreateLessonHandler   创建课堂
func CreateLessonHandler(c *gin.Context) {
	var lesson models.Lesson
	lesson.TeacherId, _ = GetCurrentUserID(c)
	var err error
	lesson.ClassId,err=strconv.ParseInt(c.Param("classid"),10,64)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,CodeInvalidParams.Msg())
	}
	// 创建帖子
	if err = mysql.CreateLesson(&lesson); err != nil {
		zap.L().Error("mysql.CreateLesson(&lesson) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)

}

// SetHomeworkHandler 布置作业
func SetHomeworkHandler(c *gin.Context) {
	homework:=c.PostForm("homework")
	lessonid,err:=strconv.ParseInt(c.Param("lessonid"),10,64)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,CodeInvalidParams.Msg())
	}

	if err = mysql.SetHomework(homework,lessonid); err != nil {
		zap.L().Error("mysql.SetHomework(homework,lessonid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// LessonOverHandler 下课
func LessonOverHandler(c *gin.Context) {
	lessonid,err:=strconv.ParseInt(c.Param("lessonid"),10,64)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,CodeInvalidParams.Msg())
	}
	if err = mysql.LessonOver(lessonid); err != nil {
		zap.L().Error("mysql.LessonOver(lessonid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)

}

// LessonSignInHandler 学生签到
func LessonSignInHandler(c *gin.Context) {
	var lessonInfo models.LessonInfo
	var err error
	lessonInfo.LessonId,err=strconv.ParseInt(c.Param("lessonid"),10,64)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,CodeInvalidParams.Msg())
	}
	lessonInfo.StudentId, _ = GetCurrentUserID(c)

	err = logic.LessonSignIn(&lessonInfo)
	if err != nil {
		zap.L().Error("logic.LessonSignIn failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy,err.Error())
		return
	}
	ResponseSuccess(c, nil)
}

// HomeworkSubmitHandler 作业提交
func HomeworkSubmitHandler(c *gin.Context) {
	lessonid,err:=strconv.ParseInt(c.Param("lessonid"),10,64)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,CodeInvalidParams.Msg())
	}
	studentid,_:= GetCurrentUserID(c)
	if err := mysql.HomeworkSubmit(lessonid,studentid); err != nil {
		zap.L().Error("mysql.HomeworkSubmit(lessonid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// LessonCheckHandler 查看课堂作业完成情况
func LessonCheckHandler(c *gin.Context) {
	lessonid,err:=strconv.ParseInt(c.Param("lessonid"),10,64)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,CodeInvalidParams.Msg())
	}
	lessoninfos,err := mysql.LessonCheck(lessonid)
	if  err != nil {
		zap.L().Error("mysql.LessonCheck(lessonid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, lessoninfos)
}

// LessonTalkRankHandler 获取课堂发言排行榜（根据页数）
func LessonTalkRankHandler(c *gin.Context) {
	lessonid:=c.Param("lessonid")
	page, err:=strconv.ParseInt(c.Param("page"),10,64)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,CodeInvalidParams.Msg())
	}

	rank, err := redis.GetTalkRank(page, lessonid)
	if err != nil {
		zap.L().Error("redis.GetTalkRank(page, lessonid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c,rank)
	//ResponseSuccess(c,lessonid)
	//ResponseSuccess(c,page)
}

// SetTestHandler 发布试题
func SetTestHandler(c *gin.Context) {

	//获取参数
	choiceNum, _ :=strconv.Atoi(c.PostForm("choiceQuestionNum"))
	fillingNum, _ :=strconv.Atoi(c.PostForm("gapFillingNum"))
	subject:=c.PostForm("subject")


	//获取登陆身份
	username,err:=GetCurrentUserName(c)

	//生成试卷
	testId, err := logic.GetTest(choiceNum,fillingNum,subject,username)
	if err != nil {
		zap.L().Error("logic.GetTest() failed", zap.Error(err))
		return
	}
	choiceJson, fillingJson, answerJson, err := redis.GetTest(testId)
	if err != nil {
		zap.L().Error("redis.GetTest(testId) failed", zap.Error(err))
		return
	}

	//s := make([]json.RawMessage, 0)
	//s = append(s, choiceJson)
	//s = append(s, fillingJson)
	//s = append(s, answerJson)

	ResponseSuccess(c,gin.H{
		"0试卷编号":testId,
		"1选择题":choiceJson,
		"2填空题":fillingJson,
		"3答案":answerJson,
	})
}

// AddChoiceQuestionHandler 向题库增加选择题
func AddChoiceQuestionHandler(c *gin.Context) {
	var q models.ChoiceQuestion
	if err := c.ShouldBind(&q); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		c.JSON(http.StatusOK, gin.H{
			"msg": RemoveTopStruct(errs.Translate(trans)),
		})
		return

	}
	err := mysql.AddChoiceQuestion(q)
	if err != nil {
		zap.L().Error("mysql.AddChoiceQuestion() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "添加成功")
}

// AddGapFillingHandler 增加填空题
func AddGapFillingHandler(c *gin.Context) {
	var q models.GapFilling
	if err := c.ShouldBind(&q); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		c.JSON(http.StatusOK, gin.H{
			"msg": RemoveTopStruct(errs.Translate(trans)),
		})
		return
	}
	err := mysql.AddGapFilling(q)
	if err != nil {
		zap.L().Error("mysql.AddGapFilling() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "添加成功")
}

// GetTestHandler 根据试卷id获取试卷题目
func GetTestHandler(c *gin.Context) {
	testId,err:=strconv.ParseInt(c.Param("testid"),10,64)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,CodeInvalidParams.Msg())
	}
	choiceJson, fillingJson, _, err := redis.GetTest(testId)
	if err != nil {
		ResponseErrorWithMsg(c,CodeServerBusy,CodeServerBusy.Msg())
		zap.L().Error("redis.GetTest(testId) failed", zap.Error(err))
		return
	}
	ResponseSuccess(c,gin.H{
		"0试卷编号":testId,
		"1选择题":choiceJson,
		"2填空题":fillingJson,
	})
}

// GetScoreHandler 提交试卷 自动批改获取分数
func GetScoreHandler(c *gin.Context)  {
	//获取testId
	testId,err:=strconv.ParseInt(c.Param("testid"),10,64)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,CodeInvalidParams.Msg())
	}
	//获取题目数量
	choiceNum,fillingNum,err:=redis.GetQuestionNum(testId)
	//获取题目答案
	var ans []models.ResponseAnswers
	//获取选择题
	{
		for i := 1; i <= choiceNum; i++ {
			ans=append(ans,models.ResponseAnswers{
				QuestionID: int64(i),
				Answer:     c.PostForm(strconv.Itoa(i )),
			})

		}
	}
	//获取填空题
	{
		for i := choiceNum+1; i <= choiceNum+fillingNum; i++ {
			ans=append(ans,models.ResponseAnswers{
				QuestionID: int64(i),
				Answer:     c.PostForm(strconv.Itoa(i)),
			})
		}
	}

	var studentScore models.StudentScore

	//比对答案获取分数
	ans,studentScore.Score,err=logic.GetTestScore(testId,ans,choiceNum,fillingNum)
	if err != nil {
		zap.L().Error("redis.GetTestScore(testId) failed", zap.Error(err))
		ResponseErrorWithMsg(c,CodeInvalidParams,err.Error())
		return
	}
	studentScore.StudentID, err =GetCurrentUserID(c)
	studentScore.StudentName,err=GetCurrentUserName(c)
	err=mysql.SetScore(testId,studentScore)
	if err != nil {
		zap.L().Error("redis.GetTestScore(testId) failed", zap.Error(err))
		ResponseErrorWithMsg(c,CodeInvalidParams,err.Error())
		return
	}

	ResponseSuccess(c,gin.H{
		"分数":studentScore.Score,
		"结果":ans,
	})
}

// GetStudentTestScoreHandler 根据testid获取该试卷学生的完成情况
func GetStudentTestScoreHandler(c *gin.Context) {
	//获取testId
	testId,err:=strconv.ParseInt(c.Param("testid"),10,64)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,CodeInvalidParams.Msg())
	}
	scores,err:=mysql.GetTestStudentScore(testId)
	if err != nil {
		zap.L().Error("mysql.GetTestStudentScore failed", zap.Error(err))
		ResponseErrorWithMsg(c,CodeServerBusy,CodeServerBusy.Msg())
		return
	}
	ResponseSuccess(c,scores)
}