package redis

import (
	"github.com/go-redis/redis"
	"go_blog/models"
	"strconv"
	"time"
)

func GetPostIDInOrder(p *models.ParamPostList) ([]string, error) {
	//依据排序方式选择指标
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(keyPostScoreZSet)
	}
	//确定查询的范围
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	//按索引从大到小查询指定范围内容
	return rdb.ZRevRange(key, start, end).Result()
}

func GetCommunityPostIDInOrder(p *models.ParamPostList) ([]string, error) {
	//依据排序方式选择指标
	oKey := getRedisKey(KeyPostTimeZSet) //排序方式键
	if p.Order == models.OrderScore {
		oKey = getRedisKey(keyPostScoreZSet)
	}
	cKey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(p.CommunityID))) //社区ID键
	//使用zinterstore取分区帖子set和帖子分数/时间zset的交集，余下逻辑相同
	key := oKey + strconv.Itoa(int(p.CommunityID)) //交集键
	if rdb.Exists(key).Val() < 1 {                 //为节省时间，仅在交集不存在的情况下进行取交集
		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX", //取交集时值的处理方式
		}, cKey, oKey)
		pipeline.Expire(key, 60*time.Second) //设置超时时间,交集的缓存存活时间为超时时间
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	//确定查询的范围
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	//按索引从大到小查询指定范围内容
	return rdb.ZRevRange(key, start, end).Result()
}

func GetPostVoteData(ids []string) (data []int64, err error) {
	pipeline := rdb.Pipeline() //通过pipeline一次查询全部的投票数，减少RTT
	for _, id := range ids {
		key := getRedisKey(keyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmds, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(ids))
	for _, cmd := range cmds {
		v := cmd.(*redis.IntCmd).Val() //断言查询结果的类型为IntCmd，返回值
		data = append(data, v)
	}

	return
}
func CreatePost(pid, cid int64) error {
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
	//设置社区
	pipeline.SAdd(getRedisKey(KeyCommunitySetPrefix+strconv.Itoa(int(cid))), pid)
	_, err := pipeline.Exec()
	return err
}
