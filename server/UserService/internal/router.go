package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Router(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "OK!",
		})
	})

	app.Post("/user/signup", func(c *fiber.Ctx) error {
		var userSignupRequest UserSignupRequest
		err := c.BodyParser(&userSignupRequest)
		if err != nil {
			panic(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		newUser := UserProfile{
			Username:  userSignupRequest.Username,
			Email:     userSignupRequest.Email,
			Password:  userSignupRequest.Password,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		UserCollection.InsertOne(ctx, newUser)
		return c.JSON(fiber.Map{
			"success": true,
			"message": "OK!",
		})
	})

	app.Post("/user/login", func(c *fiber.Ctx) error {
		var userLoginRequest UserLoginRequest
		err := c.BodyParser(&userLoginRequest)

		filter := bson.M{
			"username": userLoginRequest.Username,
		}

		var user UserProfile
		err = UserCollection.FindOne(context.TODO(), filter).Decode(&user)
		fmt.Println(user, err)
		if err != nil {
			return c.JSON(fiber.Map{
				"success": false,
				"message": "Not Ok!",
			})
		}
		token, err := createToken(userLoginRequest.Username)
		if err != nil {
			panic(err)
		}
		return c.JSON(fiber.Map{
			"success": true,
			"message": "OK!",
			"payload": token,
		})
	})
}
