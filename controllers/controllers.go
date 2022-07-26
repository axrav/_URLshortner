package controllers

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	db "Vegetaxd/Urlshortner/database"

	"github.com/gofiber/fiber/v2"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

type rq struct {
	URL string `json:"url"`
}

func randomString() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 7)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:7]
}

func Short(c *fiber.Ctx) error {
	body := new(rq)
	gen_key := randomString()
	c.BodyParser(&body)
	host := os.Getenv("host_name")
	fmt.Println(body.URL)
	fmt.Println("Without rq")
	if db.Setkey(body.URL, gen_key) {
		resp := fmt.Sprintf(`{"key" : "%s", "shortened_url" : "%s/%s"}`, gen_key, host, gen_key)
		fmt.Printf("Request Received for URL %s, Processed Successfully!", body.URL)
		return c.JSON(resp)
	} else {
		return c.Status(500).SendString("Something went wrong!")
	}

}
func Shortend(c *fiber.Ctx) error {
	key := c.Params("+")
	fmt.Println("Get rq")
	red_url := db.GetKey(key)
	fmt.Println("redirected")
	fmt.Println(red_url)
	return c.Redirect(red_url)
}
