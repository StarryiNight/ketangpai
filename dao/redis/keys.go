package redis

/*
	Redis Key
*/

const (
	KeyPostInfoHashPrefix      = "ketangpai:post:"
	KeyPostTimeZSet            = "ketangpai:post:time"
	KeyPostScoreZSet           = "ketangpai:post:score"
	KeyPostVotedZSetPrefix     = "ketangpai:post:voted:"
	KeyCommunityPostSetPrefix  = "ketangpai:community:"
	KeyTalkFrequencyZSetPrefix = "ketangpai:talk:score:"
	KeyTalkUserHashPrefix      = "ketangpai:talk:user:"
	KeyVerificationCodeHash    = "ketangpai:verificationCode"
	KeyTestChoiceHashPrefix    = "ketangpai:test:choice:"
	KeyTestAnswerHashPrefix    = "ketangpai:test:answer:"
	keyTestFillingHashPrefix   = "ketangpai:test:filling:"
	KeyChoiceNum               = "choiceNum"
	keyFillingNum              = "fillingNum"
	KeyTestChoiceNumHashPrefix = "ketangpai:test:choiceNum:"
	KeyTestFillingNumHashPrefix = "ketangpai:test:fillingNum:"
)
