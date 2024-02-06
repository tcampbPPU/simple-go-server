package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/tcampbppu/server/config"
	"github.com/tcampbppu/server/database"
	"github.com/tcampbppu/server/routes"
)

// Init Application Setup
func Init() (app *fiber.App) {
	app = fiber.New()

	// Logger
	app.Use(logger.New())

	// Load all the env variables
	config.Init(".env")

	// Connect to the database
	database.NewDatabase().Connect()

	// Middlewares
	middlewares(app)

	return routes.Register(app)
}

func middlewares(app *fiber.App) {
	app.Use(cors.New())
	app.Use(recover.New())
}
