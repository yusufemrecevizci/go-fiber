package main

import (
	"log"

	"command-line-arguments/Users/yusufemrecevizci/go-deneme/fiber-app/database/database.go"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}
