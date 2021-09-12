package redis

import (
	"encoding/json"
	"strconv"
)

func SetTest(testId int64, choiceJson []byte, fillingJson []byte, answerJson []byte) error {
	pipeline := client.TxPipeline()
	id := strconv.FormatInt(testId, 10)
	pipeline.HSet(KeyTestChoiceHashPrefix, id, choiceJson)
	pipeline.HSet(keyTestFillingHashPrefix, id, fillingJson)
	pipeline.HSet(KeyTestAnswerHashPrefix, id, answerJson)
	_, err := pipeline.Exec()
	if err != nil {
		return err
	}
	return nil
}

func GetTest(testId int64) (json.RawMessage, json.RawMessage, json.RawMessage, error) {
	id := strconv.FormatInt(testId, 10)
	choiceJson := json.RawMessage(client.HGet(KeyTestChoiceHashPrefix, id).Val())

	fillingJson := json.RawMessage(client.HGet(keyTestFillingHashPrefix, id).Val())

	answerJson := json.RawMessage(client.HGet(KeyTestAnswerHashPrefix, id).Val())
	return choiceJson, fillingJson, answerJson, nil
}

func GetAnswer(testId int64) ([]byte, error) {
	str := client.HGet(KeyTestAnswerHashPrefix, strconv.FormatInt(testId, 10)).Val()
	return []byte(str), nil
}

func SetQuestionNum(testId int64, choiceNum int, fillingNum int) error {
	id := strconv.FormatInt(testId, 10)
	pipeline := client.Pipeline()
	pipeline.HSet(KeyTestChoiceNumHashPrefix, id, choiceNum)
	pipeline.HSet(KeyTestFillingNumHashPrefix, id, fillingNum)
	_, err := pipeline.Exec()
	if err != nil {
		return err
	}
	return nil

}

func GetQuestionNum(testId int64) (choiceNum int, fillingNum int, err error) {
	id := strconv.FormatInt(testId, 10)
	choiceNum, err =strconv.Atoi(client.HGet(KeyTestChoiceNumHashPrefix,id).Val())
	if err != nil {
		return 0, 0, err
	}
	fillingNum, err =strconv.Atoi(client.HGet(KeyTestFillingNumHashPrefix,id).Val())
	if err != nil {
		return 0, 0, err
	}
	return
}


