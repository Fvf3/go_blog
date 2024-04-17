package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"go_blog/settings"
)

var rdb *redis.Client

func Init(config *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			config.Host,
			config.Port,
		),
		Password: config.Password,
		DB:       config.Db,
		PoolSize: config.PoolSize,
	})
	_, err = rdb.Ping().Result()
	return
}

func Close() {
	_ = rdb.Close()
}
