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

	// TODO: Middleware these suckas for user auth policy
	poll := api.Group("poll")
	poll.Get("/", controllers.ListPolls)
	poll.Get("/:id", controllers.GetPoll)
	poll.Post("/", controllers.CreatePoll)
	poll.Put("/:id", controllers.UpdatePoll)
	poll.Delete("/:id", controllers.DeletePoll)
	poll.Post("/:id/vote", controllers.RegisterVote)
}
