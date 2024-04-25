package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"time"
)

const (
	oneWeekSeconds = 7 * 24 * 60 * 60
	scorePerVote   = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已结束")
)

func VoteForPost(uid, pid string, direction float64) error { //由于ZScore方法键为string，返回值为float
	//1.判断是否在投票时间内
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), pid).Val() //从redis获取发帖时间
	if float64(time.Now().Unix())-postTime > oneWeekSeconds {
		return ErrVoteTimeExpire
	}
	//2.更新帖子分数
	old := rdb.ZScore(getRedisKey(keyPostVotedZSetPrefix+pid), uid).Val()
	diff := (direction - old) * scorePerVote
	//对分数和投票用户记录的操作应当在一个pipeline完成
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(keyPostScoreZSet), diff, pid)
	//3.更新投票用户记录
	pipeline.ZAdd(getRedisKey(keyPostVotedZSetPrefix+pid), redis.Z{
		Score:  direction,
		Member: uid,
	})
	_, err := pipeline.Exec()
	return err
}

func CreatePost(pid int64) error {
	pipeline := rdb.TxPipeline() //我们希望记录发帖时间和设置初始分数的事务同时发生，因此用pipeline
	//记录发帖时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: pid,
	})
	//设置初始分数
	pipeline.ZAdd(getRedisKey(keyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: pid,
	})
	_, err := pipeline.Exec()
	return err
}
