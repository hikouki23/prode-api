package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/test", func (c fiber.Ctx)  {
		c.SendString("funca")
	})

	log.Fatal(app.Listen(":3000"))
}
