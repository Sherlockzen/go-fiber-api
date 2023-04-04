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

	app.Listen("0.0.0.0:27027")
}
