package helpers

import (
	"Vegetaxd/Urlshortner/database"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

// Generating  a random string
func RandomString() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 7)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:7]
}

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
