package config

import (
	"coco-server/util/infras"

	"github.com/go-redis/redis/v7"
)

var (
	KV *redis.Client
)

func initRedis() {
	KV = redis.NewClient(&redis.Options{
		Addr:       Conf.Redis.Addr,
		Password:   Conf.Redis.Pwd,
		DB:         0,
		PoolSize:   30,
		MaxRetries: 3,
	})
	err := KV.Ping().Err()
	infras.Throw(err)
}
