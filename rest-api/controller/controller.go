package controller

import (
	database "github.com/Lefree111/go-fiber-rest-api/api/entity"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var (
	datas = map[string]database.Data{}
)

func SetupRoute() *fiber.App {
	app := *fiber.New()
	app.Post("/api/v1/create", createData)
	app.Get("/api/v1/getapi/id", readApi)
	app.Get("/api/v1/getapis/", ReadApis)
	app.Put("/api/v1/update/:id", updateApi)
	app.Delete("/api/v1/delete/:id", deleteApi)
	return &app
}

func createData(c *fiber.Ctx) error {
	data := new(database.Data)

	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	data.Id = uuid.New().String()
	datas[data.Id] = *data

	c.Status(200).JSON(&fiber.Map{
		"Data": data,
	})
	return nil
}

func readApi(c *fiber.Ctx) error {
	id := c.Params("id")

	if readApi, ok := datas[id]; ok {
		c.Status(200).JSON(&fiber.Map{
			"data": readApi,
		})
	} else {
		c.Status(404).JSON(&fiber.Map{
			"error": "data not found",
		})
	}
	return nil
}

func ReadApis(c *fiber.Ctx) error {
	c.Status(200).JSON(&fiber.Map{
		"data": datas,
	})
	return nil
}

func updateApi(c *fiber.Ctx) error {
	updateApi := new(database.Data)

	err := c.BodyParser(updateApi)

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	id := c.Params("id")
	if data, ok := datas[id]; ok {
		data.User_id = updateApi.User_id
		data.Title = updateApi.Title
		data.Body = updateApi.Body
		datas[id] = data
		c.Status(200).JSON(&fiber.Map{
			"datas": data,
		})
	} else {
		c.Status(404).JSON(&fiber.Map{
			"error": "data not found",
		})

	}
	return nil
}

func deleteApi(c *fiber.Ctx) error {
	id := c.Params("id")

	if _, ok := datas[id]; ok {
		delete(datas, id)
		c.Status(200).JSON(&fiber.Map{
			"message": "data deleted successfully",
		})
	} else {
		c.Status(404).JSON(&fiber.Map{
			"error": "data not found",
		})
	}

	return nil
}
