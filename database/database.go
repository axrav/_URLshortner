package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func Newdb() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_db"),
		Username: os.Getenv("redis_username"),
		Password: os.Getenv("redis_password"),
		DB:       0})
	return client
}
