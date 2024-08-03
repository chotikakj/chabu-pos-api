package ReportRoutes

import (
	ReportController "pos-api/api/report/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	app := router.Group("report")
	app.Post("/get-home-report", ReportController.GetHomeReport)
}
