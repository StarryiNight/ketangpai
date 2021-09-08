package redis

import (
	"fmt"
	"go.uber.org/zap"
)

const (
	TalkScore  = 1
	TalkPerAge = 10
)

type UserScore struct {
	UserName string  `json:"username"`
	UserID   string  `json:"userid"`
	Score    int `json:"score"`
}


// TalkFrequency 发言频率保存到redis
func TalkFrequency(userID string,userName string,lessonID string) (err error) {
	//将发言的用户id和次数存入redis
	key := KeyTalkFrequencyZSetPrefix+lessonID
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(key, TalkScore, userID)

	//将发言的用户id和对应的姓名存入hash
	key=KeyTalkUserHashPrefix+lessonID
	pipeline.HSet(key,userID,userName)

	//开始事务
	_, err = pipeline.Exec()
	return
}

// GetTalkRank 获取课堂发言排行榜的第page页
func GetTalkRank(page int64,lessonID string) (users []UserScore ,err error) {
	key := KeyTalkFrequencyZSetPrefix+lessonID
	start := (page - 1) * TalkPerAge
	end := start + PostPerAge - 1

	result, err :=client.ZRevRangeWithScores(key,start,end).Result()
	if err != nil {
		zap.L().Error("client.ZRevRange(key,start,end).Result() failed!",zap.Error(err))
		return nil,err
	}

	users=make([]UserScore,len(result))
	key=KeyTalkUserHashPrefix+lessonID
	var i int = 0
	for _,z:=range result{
		users[i].UserID=fmt.Sprintf("%v",z.Member)
		users[i].Score= int(z.Score)
		users[i].UserName=client.HGet(key,users[i].UserID).Val()
		i++
	}
	return users,nil
}
