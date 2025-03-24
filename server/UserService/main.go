package main

import (
	"Medium-UserService/internal"

	"github.com/gofiber/fiber/v2"
)

func main() {
	internal.InitMongoDB()
	app := fiber.New()

	internal.Router(app)

	app.Listen("3000")
}
