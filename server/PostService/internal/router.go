package internal

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "OK!",
		})
	})

	app.Post("/post/create", func(c *fiber.Ctx) error {
		var postCreateRequest PostCreateRequest
		err := c.BodyParser(&postCreateRequest)
		if err != nil {
			return c.JSON(fiber.Map{
				"success": false,
				"message": "NotOK!",
			})
		}

		newPost := PostSchema{
			Title:     postCreateRequest.Title,
			Content:   postCreateRequest.Content,
			Category:  postCreateRequest.Category,
			LikeCount: 0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		PostCollection.InsertOne(context.TODO(), newPost)

		return c.JSON(fiber.Map{
			"success": true,
			"message": "OK!",
		})
	})

	app.Post("/post/delete", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"success": true,
			"message": "OK!",
		})
	})
}
