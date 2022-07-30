package routers

import (
	"Vegetaxd/Urlshortner/controllers"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Routes() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization ",
	}))
	app.Post("/short", controllers.Short)
	app.Static("/", "./static")
	app.Get("/:key", controllers.Redirectit)
	var port string
	if os.Getenv("port") == "" {
		port = ":8000"
	} else {
		port = fmt.Sprintf(":%s", os.Getenv("port"))
	}
	fmt.Println(port)
	app.Listen(port)
	return app
}
