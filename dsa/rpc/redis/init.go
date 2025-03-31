package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/study825/director/dsa/config"
)

var RedisClient *redis.Client

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: config.GetConf().RedisConf.Address,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		//panic(err)
	}
}
