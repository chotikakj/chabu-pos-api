package router

import (
	BillingRoutes "pos-api/api/billing"
	ReportRoutes "pos-api/api/report"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to iampos api.",
		})
	})
	api := app.Group("api", logger.New())
	BillingRoutes.SetupRoutes(api)
	ReportRoutes.SetupRoutes(api)
}
