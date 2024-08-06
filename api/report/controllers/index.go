package ReportController

import (
	Schema "pos-api/api/report/schema"
	ReportService "pos-api/api/report/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New()

func GetHomeReport(c *fiber.Ctx) error {
	body := new(Schema.ReportHomeDto)
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
	result, err := ReportService.GetHomeReport(*body)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "Error",
			"message": err.Error(),
		})
	}
	if result == nil {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"status": "NotFound",
			"result": nil,
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "success",
		"result": result,
	})
}

func GetBillDetail(c *fiber.Ctx) error {
	result, err := ReportService.GetBillDetail(c.Params("bill_id"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "Error",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "success",
		"result": result,
	})
}

func GetBillTransaction(c *fiber.Ctx) error {
	body := new(Schema.ReportHomeDto)
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
	result, err := ReportService.GetBillTransaction(*body)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "Error",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "success",
		"result": result,
	})
}
