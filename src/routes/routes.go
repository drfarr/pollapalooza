package routes

import (
	"pollapalooza/src/controllers"

	"pollapalooza/src/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {
	app.Use(middlewares.Security)

	api := app.Group("api", logger.New())

	auth := api.Group("auth")
	auth.Post("signup", controllers.SignUp)
	auth.Post("signin", controllers.SignIn)
	auth.Get("signout", controllers.SignOut)
	auth.Get("user", controllers.User)

	// TODO: Middleware these suckas for user auth policy
	poll := api.Group("poll")
	pollAuthenticated := poll.Use(middlewares.IsAuthenticated)
	poll.Get("/:id", controllers.GetPoll)
	poll.Post("/:id/vote", controllers.RegisterVote)
	pollAuthenticated.Get("/", controllers.ListPolls)
	pollAuthenticated.Post("/", controllers.CreatePoll)
	pollAuthenticated.Put("/:id", controllers.UpdatePoll)
	pollAuthenticated.Delete("/:id", controllers.DeletePoll)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}
