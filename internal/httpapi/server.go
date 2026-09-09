package httpapi

import (
	"github.com/beparykamrul-dev/cgta/internal/config"
	"github.com/gofiber/fiber/v2"
)

func New(cfg config.Config) *fiber.App {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
	app.Get("/ready", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ready"})
	})
	return app
}
