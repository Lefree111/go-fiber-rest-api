package main

import (
	database "app/api/entity"

	"app/rest-api/controller"

	"github.com/gofiber/fiber/v2"
)

func init() {
	database.NewPostgreSQLClient()
}

func main() {

	app := controller.SetupRoute()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Go Fiber API")
	})

	app.Listen(":5000")
}
