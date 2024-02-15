package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tcampbppu/server/app/models"
	"github.com/tcampbppu/server/app/traits/harmony"
	"github.com/tcampbppu/server/app/traits/paginate"
	"github.com/tcampbppu/server/database"
)

func Register(app *fiber.App) *fiber.App {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/api", func(c *fiber.Ctx) error {
		var products []models.Product
		database.DB.Find(&products)

		return harmony.Success(c, "Success", products)
	})

	app.Get("/paginate", func(c *fiber.Ctx) error {
		return paginate.Paginate(c, database.DB, models.Product{})
	})

	app.Get("/cursor", func(c *fiber.Ctx) error {
		return paginate.CursorPaginate(c, database.DB)
	})

	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("Panic!")
	})

	return app
}
