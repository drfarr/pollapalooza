package controllers

import (
	"fmt"
	"pollapalooza/src/database"
	"pollapalooza/src/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var key = []byte("SecretYouShouldHide")

type UserClaim struct {
	jwt.RegisteredClaims
	ID    uint
	Email string
	Name  string
}

func generateToken(id uint, email string, name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{},
		ID:               id,
		Email:            email,
		Name:             name,
	})

	signedString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", fmt.Errorf("error creating signed string: %v", err)
	}
	return signedString, nil
}

/***
*** Create a User account.
***/

func Signup(c *fiber.Ctx) error {
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

	return c.JSON(user)
}

/***
*** Authenticate a user with email and password.
***/

func Signin(c *fiber.Ctx) error {
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

	return c.JSON(fiber.Map{
		"message": "Success",
	})

}

/***
*** Get a user from a cookie, the return the user data from the database
***/
func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	var userClaim UserClaim
	token, err := jwt.ParseWithClaims(cookie, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil || !token.Valid {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	var user models.User
	database.DB.Where("id = ?", userClaim.ID).First(&user)

	return c.JSON(user)
}
