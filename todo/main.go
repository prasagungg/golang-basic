package main

import "github.com/gofiber/fiber/v3"

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
    StrictRouting: true,
    ServerHeader:  "Fiber",
    AppName: "Test App v1.0.1",
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
