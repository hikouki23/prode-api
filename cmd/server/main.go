package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"hikouki.com/prode/internal/config"
	"hikouki.com/prode/internal/storage"
)

func main() {
	cfg := config.LoadConfig()
	app := fiber.New()

	app.Get("/test", func(c fiber.Ctx) {
		c.SendString("funca")
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.Port)))

	db, err := storage.New(cfg)
}
