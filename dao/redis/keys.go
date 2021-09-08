package redis

/*
	Redis Key
*/

const (
	KeyPostInfoHashPrefix  = "ketangpai:post:"
	KeyPostTimeZSet        = "ketangpai:post:time"
	KeyPostScoreZSet       = "ketangpai:post:score"
	KeyPostVotedZSetPrefix = "ketangpai:post:voted:"
	KeyCommunityPostSetPrefix = "ketangpai:community:"
	KeyTalkFrequencyZSetPrefix      = "ketangpai:talk:score:"
	KeyTalkUserHashPrefix ="ketangpai:talk:user:"
)
