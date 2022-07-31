package main

import (
	"Vegetaxd/Urlshortner/routers"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println("[INFO] .env Loaded Successfully!")
	fmt.Println("[INFO] Server has started!")
	routers.Routes()
}
