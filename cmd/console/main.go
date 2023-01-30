package main

import (
	database "github.com/Lefree111/go-fiber-rest-api/api/entity"

	"github.com/Lefree111/go-fiber-rest-api/rest-api/controller"

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
