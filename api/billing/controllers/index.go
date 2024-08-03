package BillingController

import (
	Schema "pos-api/api/billing/schema"
	BillingService "pos-api/api/billing/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New()

func CreateBill(c *fiber.Ctx) error {
	body := new(Schema.CreateBillDto)
	c.BodyParser(&body)
	err := Validator.Struct(body)
	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Field()+"is "+err.Tag())
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "validate error",
			"message": errors,
		})
	}
	err = BillingService.CreateBill(*body)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "Error",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "success",
		"result": "Create bill successfully.",
	})
}
