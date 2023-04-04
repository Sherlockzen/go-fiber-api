package main

import (
	"go-fiber-api/configs"
	"go-fiber-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app)

	app.Listen("1.1.14.1:27027")
}
