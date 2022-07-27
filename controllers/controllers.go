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

var host string = os.Getenv("host_name")

func Short(c *fiber.Ctx) error {
	body := new(rq)
	gen_key := randomString()
	c.BodyParser(&body)
	fmt.Println(body.URL)
	fmt.Println("Without rq")
	if db.Setkey(body.URL, gen_key) {
		resp := fmt.Sprintf("%s/%s", host, gen_key)
		fmt.Printf("Request Received for URL %s, Processed Successfully!", body.URL)
		return c.JSON(fiber.Map{"key": gen_key, "short_url": resp})
	} else {
		return c.Status(500).JSON(fiber.Map{"error": "Cannot Parse JSON, something went wrong"})
	}

}
func Redirectit(c *fiber.Ctx) error {
	key := c.Params("key")
	fmt.Println(key)
	red_url := db.GetKey(key)
	if red_url == "No Key found" {
		return c.Status(500).JSON(fiber.Map{"error": "Key not found"})
	}
	fmt.Println(red_url)
	return c.Redirect(red_url)

}
