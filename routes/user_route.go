package routes

import (
	"go-fiber-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	//All routes related to users comes here
	app.Post("/user", controllers.CreateUser)
}
