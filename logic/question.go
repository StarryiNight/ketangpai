package logic

import (
	"encoding/json"
	"go.uber.org/zap"
	"ketangpai/dao/mysql"
	"ketangpai/dao/redis"
	"ketangpai/models"
	"ketangpai/pkg/snowflake"
)

func GetTest(ChoiceNum int, fillingNum int,subject string,publisher string) (testId int64,err error) {
	var test models.Test
	test.ChoiceQuestion, err = mysql.GetChoiceQuestion(ChoiceNum,subject)
	if err != nil {
		return 0,err
	}
	test.GapFilling, err = mysql.GetGapFilling(fillingNum,subject)
	if err != nil {
		return 0,err
	}
	test.TestID, err = snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
	}

	err = TestMarshalAndSet(test)
	if err != nil {
		zap.L().Error("TestMarshalAndSet failed", zap.Error(err))
		return 0,err
	}
	err = redis.SetQuestionNum(test.TestID,ChoiceNum,fillingNum)
	if err != nil {
		zap.L().Error("redis.SetQuestionNum failed", zap.Error(err))
		return 0,err
	}
	//插入数据库
	err = mysql.SetTest(test.TestID,subject,publisher)
	if err != nil {
		zap.L().Error("mysql.SetTest() failed", zap.Error(err))
		return 0, err
	}

	return test.TestID,err
}

func TestMarshalAndSet(test models.Test) error {
	var choiceQuestion []models.ChoiceQuestionWithoutAnswer
	var answer []models.Answers
	var fillingQuestion []models.GapFillingWithoutAnswer
	count := 0
	for _, i := range test.ChoiceQuestion {
		choiceQuestion=append(choiceQuestion,models.ChoiceQuestionWithoutAnswer{
			QuestionID: int64(count + 1),
			Content:    i.Content,
			Options:    i.Options,
		})
		answer=append(answer,models.Answers{
			QuestionID: int64(count + 1),
			Answer:     i.Answer,
		})
		count++
	}
	for _, i := range test.GapFilling {
		fillingQuestion=append(fillingQuestion,models.GapFillingWithoutAnswer{
			QuestionID: int64(count + 1),
			Content:    i.Content,
		})
		answer=append(answer,models.Answers{
			QuestionID: int64(count + 1),
			Answer:     i.Answer,
		})
		count++
	}
	choiceJson, err := json.Marshal(choiceQuestion)
	if err != nil {
		return err
	}

	fillingJson, err :=  json.Marshal(fillingQuestion)
	if err != nil {
		return err
	}
	answerJson, err :=  json.Marshal(answer)
	if err != nil {
		return err
	}

	err = redis.SetTest(test.TestID, choiceJson, fillingJson, answerJson)
	if err != nil {
		return err
	}
	return nil
}

