package database

import "github.com/go-redis/redis/v9"

func InitRedis() *redis.Client {

	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
