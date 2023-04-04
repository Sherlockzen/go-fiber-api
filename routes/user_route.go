package routes

import (
	"go-fiber-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	//Post a new user
	app.Post("/user", controllers.CreateUser)

	//Get user by a ID
	app.Get("/user/:userId", controllers.GetAUser)

	//Update a user
	app.Put("/user/:userId", controllers.EditAUser)

	//delete a user
	app.Delete("/user/:userId", controllers.DeleteAUser)
}
