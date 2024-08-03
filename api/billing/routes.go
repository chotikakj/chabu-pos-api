package BillingRoutes

import (
	BillingController "pos-api/api/billing/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	app := router.Group("billing")
	app.Post("/create-transaction", BillingController.CreateBill)
}
