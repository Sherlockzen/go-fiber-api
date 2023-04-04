package main

import (
	"go-fiber-api/configs"
	"go-fiber-api/routes"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app)

	//start server
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	app.Listen(":" + port)
}
