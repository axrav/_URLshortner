package db

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

func Setkey(url string, key string) bool {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_db"),
		Username: os.Getenv("redis_username"),
		Password: os.Getenv("redis_password"),
		DB:       0,
	})
	err := client.Set(ctx, key, url, 0).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func GetKey(Key string) string {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_db"),
		Username: os.Getenv("redis_username"),
		Password: os.Getenv("redis_password"),
		DB:       0,
	})
	ctx := context.Background()
	val, err := client.Get(ctx, Key).Result()
	if err != nil {
		panic(err)
	}
	return val
}
