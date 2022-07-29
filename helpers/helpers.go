package helpers

import (
	"Vegetaxd/Urlshortner/database"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func Setkey(url string, key string) bool {

	client := database.Newdb()
	defer client.Close()
	ctx := database.Ctx
	err := client.Set(ctx, key, url, 0).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func GetKey(Key string) string {
	client := database.Newdb()
	defer client.Close()
	ctx := database.Ctx
	val, err := client.Get(ctx, Key).Result()
	if err == redis.Nil {
		return "No Key found"
	}
	return val
}
