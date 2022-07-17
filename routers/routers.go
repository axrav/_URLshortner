package routers

import (
	"Vegetaxd/Urlshortner/controllers"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Routes() *fiber.App {
	app := fiber.New()
	app.Post("/short", controllers.Short)
	app.Static("/", "./static")
	app.Get("/+", controllers.Shortend)
        var port string
        if os.Getenv("port") == ""{
		port = ":8000"
        } else {
	        port = fmt.Sprintf(":%s", os.Getenv("port"))}
	fmt.Println(port)
	app.Listen(port)
	return app
}
