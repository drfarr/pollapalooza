package controllers

import (
	"pollapalooza/src/database"
	"pollapalooza/src/middlewares"
	"pollapalooza/src/models"

	"github.com/gofiber/fiber/v2"
)

/***
*** Update user info
***/
func UpdateUser(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id, err := middlewares.GetUserId(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	user := models.User{
		Id:        id,
		FirstName: data["firstName"],
		LastName:  data["lastName"],
		Email:     data["email"],
	}

	database.DB.Model(&user).Updates(&user)
	return c.Status(fiber.StatusOK).JSON(user)
}

/***
*** Update user password
***/
func UpdateUserPassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["passwordConfirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	id, err := middlewares.GetUserId(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	user := models.User{
		Id: id,
	}

	user.SetPassword(data["password"])

	database.DB.Model(&user).Updates(&user)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}
