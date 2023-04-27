package routes

import (
	"pollapalooza/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	admin := api.Group("auth")
	admin.Post("signup", controllers.Signup)
	admin.Post("signin", controllers.Signin)
	admin.Get("user", controllers.User)
}
