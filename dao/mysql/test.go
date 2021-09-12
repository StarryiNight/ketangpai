package mysql

import (
	"go.uber.org/zap"
	"ketangpai/models"
	"ketangpai/pkg/snowflake"
	"time"
)

// GetChoiceQuestion 获取随机的num个选择题
func GetChoiceQuestion(num int, subject string) ([]models.ChoiceQuestion, error) {
	sqlStr := `	SELECT question_id, type, question_id, content, options, answer 
				FROM (select * FROM choicequestion WHERE type=?) as new
				ORDER BY RAND() LIMIT ?`
	var questions []models.ChoiceQuestion
	err := db.Select(&questions, sqlStr, subject, num)
	if err != nil {
		zap.L().Error("select  choicequestion failed", zap.Error(err))
		return nil, err
	}

	return questions, nil
}

// GetGapFilling 获取随机的num个填空题
func GetGapFilling(num int, subject string) ([]models.GapFilling, error) {
	sqlStr := `	SELECT type, question_id, content, answer
				FROM (select * FROM gapfilling WHERE type=?) as new
				ORDER BY RAND() LIMIT ?`
	var questions []models.GapFilling
	err := db.Select(&questions, sqlStr, subject, num)
	if err != nil {
		zap.L().Error("select  choicequestion failed", zap.Error(err))
		return nil, err
	}
	return questions, nil
}

// AddGapFilling 添加填空题
func AddGapFilling(question models.GapFilling) (err error) {
	question.QuestionID, err = snowflake.GetID()
	if err != nil {
		return err
	}
	sqlStr := `insert into gapfilling(
	question_id, content, answer,type)
	values(?,?,?,?)`
	_, err = db.Exec(sqlStr, question.QuestionID, question.Content, question.Answer, question.Type)
	if err != nil {
		return err
	}
	return nil
}

// AddChoiceQuestion 添加选择题
func AddChoiceQuestion(question models.ChoiceQuestion) (err error) {
	question.QuestionID, err = snowflake.GetID()
	sqlStr := `insert into choicequestion(
	question_id, content, options,answer,type)
	values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, question.QuestionID, question.Content, question.Options, question.Answer, question.Type)
	if err != nil {
		return err
	}
	return nil
}

// SetTest 发布试卷
func SetTest(testId int64, subject string, publisher string) (err error) {
	sqlStr := `insert into test(test_id, type, publisher, creat_time)
	values(?,?,?,?)`
	_, err = db.Exec(sqlStr, testId, subject, publisher, time.Now())
	if err != nil {
		return err
	}
	return nil
}

// SetScore 在test——score表插入成绩
func SetScore(testId int64, studentScore models.StudentScore) (err error) {
	sqlStr := `insert into test_score(test_id, student_id, student_name, submit_time, score)
	values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, testId, studentScore.StudentID, studentScore.StudentName, time.Now(), studentScore.Score)
	if err != nil {
		return err
	}
	return nil
}

// GetTestStudentScore  查看试卷完成情况
func GetTestStudentScore(testId int64) ([]models.StudentScore, error) {
	var scores []models.StudentScore
	sqlStr := `select  student_id, student_name, submit_time, score from test_score where test_id=? order by score`
	err := db.Select(&scores, sqlStr, testId)
	if err != nil {
		return nil, err
	}
	return scores, nil
}
