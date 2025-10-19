package connector

import (
	"context"
	"github.com/hugo8680/goat/framework/config"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var cache *redis.Client

func ConnectToRedis() {
	conf := config.GetSetting()
	cache = redis.NewClient(&redis.Options{
		Addr:     conf.Cache.Host + ":" + strconv.Itoa(conf.Cache.Port),
		Password: conf.Cache.Password,
		DB:       conf.Cache.Database,
	})

	_, err := cache.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

func GetCache() *redis.Client {
	return cache
}
