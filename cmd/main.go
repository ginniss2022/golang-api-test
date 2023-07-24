package main

import (
	"log"

	"github.com/ginniss2022/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Post("/api/products", handlers.CreateProduct)
	log.Fatal(app.Listen(":3000"))
}
