package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tcampbppu/server/app/models"
	"github.com/tcampbppu/server/database"
)

func Register(app *fiber.App) *fiber.App {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/api", func(c *fiber.Ctx) error {

		var products []models.Product
		database.DB.Find(&products)

		return c.
			Status(fiber.StatusOK).
			JSON(&fiber.Map{
				"ip":      c.IP(),
				"headers": c.Get("User-Agent"),
				"success": true,
				"data":    products,
			})
	})

	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("Panic!")
	})

	return app
}
