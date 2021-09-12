package logic

import (
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"ketangpai/dao/mysql"
	"ketangpai/dao/redis"
	"ketangpai/models"
	"sort"
	"strings"
	"time"
)

const FullMark= 100.0

func LessonSignIn(lessonInfo *models.LessonInfo) (err error) {
	lesson, err := mysql.GetLessonById(lessonInfo.LessonId)
	if err != nil {
		zap.L().Error("mysql.GetLessonById(lessonInfo.LessonId) failed", zap.Int64("lessonId", lessonInfo.LessonId), zap.Error(err))
		return  err
	}
	//防止重复签到
	flag,err:=mysql.CheckSignIn(lessonInfo)
	if flag {
		return errors.New("不能重复签到")
	}

	lessonInfo.SignInTime=time.Now()
	lessonInfo.SignInStatus=1
	//u, _ := time.ParseInLocation("2006-01-02 15:04:05", "2030-01-01 00:00:00",time.Local)
	if !lesson.EndTime.IsZero(){
		if lessonInfo.SignInTime.After(lesson.EndTime){
			lessonInfo.SignInStatus=-1
		}
	}
	err = mysql.LessonSignIn(lessonInfo)
	if err != nil {
		zap.L().Error(" mysql.LessonSignIn(lessonInfo) failed", zap.Error(err))
		return err
	}
	return
}

func GetTestScore(testId int64,postAnswer []models.ResponseAnswers,choiceNum int,fillingNum int) (result []models.ResponseAnswers ,score float32, err error) {
	//从redis中获取answer数组
	data, err :=redis.GetAnswer(testId)
	if err != nil {
		zap.L().Error("redis.GetAnswer(testId) failed", zap.Error(err))
		return
	}
	var ans []models.Answers
	err=json.Unmarshal(data,&ans)
	if err != nil {
		zap.L().Error("json.Unmarshal(data,&ans) failed", zap.Error(err))
		return
	}
	count:=0
	//比对选择题答案
	for i := 0; i < choiceNum; i++ {
		if ans[i].QuestionID!=postAnswer[i].QuestionID{
			return nil, 0,errors.New("答案顺序不匹配")
		}
		if ans[i].Answer == SortString(postAnswer[i].Answer) {
			postAnswer[i].Result=true
			count++
		}else{
			postAnswer[i].Result=false
		}
	}
	//比对填空题答案
	for i := choiceNum; i < choiceNum+fillingNum; i++ {
		if ans[i].QuestionID!=postAnswer[i].QuestionID{
			return nil, 0,errors.New("答案顺序不匹配")
		}
		if ans[i].Answer == (postAnswer[i].Answer) {
			postAnswer[i].Result=true
			count++
		}else{
			postAnswer[i].Result=false
		}
	}
	score= FullMark / float32(fillingNum + choiceNum)*float32(count)

	//保存到数据库中


	return postAnswer,score,nil
}


// SortString 将string转为大写并排序
func SortString(w string) string {
	w = strings.ToUpper(w)
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}