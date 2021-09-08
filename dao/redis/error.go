package redis

import "errors"

var (
	ErrorVoteTimeExpire = errors.New("已过投票时间")
	ErrorVoted          = errors.New("已经投过票了")
	ErrorSaveFailed     = errors.New("redis保存失败")
)
