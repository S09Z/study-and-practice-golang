package helloworldme

import (
	fiber "github.com/gofiber/fiber/v2"
)

func httpHandle() {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Listen(":8088")
}
