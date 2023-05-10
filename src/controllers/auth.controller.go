package controllers

import (
	"fmt"
	"os"
	"pollapalooza/src/database"
	"pollapalooza/src/middlewares"
	"pollapalooza/src/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt/v4"
)

var key = []byte(os.Getenv("JWT_SECRET"))

var store = session.New()

func generateToken(id uint, email string, name string) (string, error) {

	payload := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(int(id)),
		ExpiresAt: &jwt.NumericDate{time.Now().Add(time.Hour * 24)},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedString, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("error creating signed string: %v", err)
	}
	return signedString, nil
}

/***
*** Create a User account.
***/

func SignUp(c *fiber.Ctx) error {
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

	user := models.User{
		FirstName: data["firstName"],
		LastName:  data["lastName"],
		Email:     data["email"],
	}

	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.Status(fiber.StatusOK).JSON(user)
}

/***
*** Authenticate a user with email and password.
***/

func SignIn(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	token, err := generateToken(user.Id, user.Email, user.LastName)
	if err != nil {
		return err
	}
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Wrong password",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}
	if err := sess.Save(); err != nil {
		panic(err)
	}

	return c.Status(fiber.StatusOK).JSON(
		user,
	)

}

/***
*** Logout a user
***/

func SignOut(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

/***
*** Get a user from a cookie, the return the user data from the database
***/
func Me(c *fiber.Ctx) error {
	id, err := middlewares.GetUserId(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	return c.Status(fiber.StatusOK).JSON(user)
}
