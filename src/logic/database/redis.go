package database

import (
	"github.com/go-redis/redis"
	"time"
)

var redisClient map[string]*redis.Client

func GetClient(shard string) *redis.Client {
	client := redisClient[shard]
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:         "127.0.0.1:6379",
			DialTimeout:  time.Second,
			ReadTimeout:  time.Second,
			WriteTimeout: time.Second,
		})
		redisClient[shard] = client
	}
	return client
}

func GetDefaultClient() *redis.Client {
	return GetClient("default")
}
