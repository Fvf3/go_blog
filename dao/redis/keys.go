package redis

// redis中的key
const (
	KeyPrefix              = "blog:"       //命名空间
	KeyPostTimeZSet        = "post:time"   // 帖子ID 发帖时间
	keyPostScoreZSet       = "post:score"  //帖子ID 投票分数
	keyPostVotedZSetPrefix = "post:voted:" //用户ID 投票类型; 参数为post id
	KeyCommunitySetPrefix  = "community:"  //每个分区中的帖子ID
)

// 获取当前命名空间的key
func getRedisKey(key string) string {
	return KeyPrefix + key
}
