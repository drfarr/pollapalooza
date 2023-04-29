package controllers

import (
	"pollapalooza/src/database"
	"pollapalooza/src/models"

	"github.com/gofiber/fiber/v2"
)

/**
 * `GET /polls` - retrieves all the available polls
 */
func ListPolls(c *fiber.Ctx) error {
	// TODO: Make it work
	poll := []models.Poll{}
	database.DB.Find(&poll)

	if err := c.BodyParser(poll); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(poll)
}

/**
 * `POST /polls` - creates a new poll
 **/
func CreatePoll(c *fiber.Ctx) error {
	// TODO: Make it work
	poll := []models.Poll{}
	if err := c.BodyParser(poll); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Create(&poll)

	return c.Status(200).JSON(poll)
}

/**
* `DELETE /polls/:id` - deletes an existing poll by its ID
**/

func DeletePoll(c *fiber.Ctx) error {
	// TODO: Make it work
	poll := []models.Poll{}
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	return c.Status(200).JSON(poll)
}

/**
* `PUT /polls/:id` - updates an existing poll by its ID
**/
func UpdatePoll(c *fiber.Ctx) error {
	// TODO: Make it work
	poll := []models.Poll{}
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	return c.Status(200).JSON(poll)
}

/**
* `GET /polls/:id` - retrieves a specific poll by its ID
**/
func GetPoll(c *fiber.Ctx) error {
	// TODO: Make it work
	poll := []models.Poll{}
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	return c.Status(200).JSON(poll)
}

/**
* `POST /polls/:id/vote`
**/
func RegisterVote(c *fiber.Ctx) error {
	// TODO: Make it work
	poll := []models.Poll{}
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	return c.Status(200).JSON(poll)
}
