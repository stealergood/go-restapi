package main

import (
	"github.com/gofiber/fiber/v2"
	"go-restapi/controllers/bookcontroller"
	"go-restapi/models"
)

func main() {
	models.ConnectDB()

	app := fiber.New()

	api := app.Group("/api")
	book := api.Group("/books")

	book.Get("/", bookcontroller.Index)
	book.Get("/:id", bookcontroller.Show)
	book.Post("/", bookcontroller.Create)
	book.Put("/:id", bookcontroller.Update)
	book.Delete("/:id", bookcontroller.Delete)

	app.Listen(":5000")
}
