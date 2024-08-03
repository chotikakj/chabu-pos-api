package middleware

import (
	"pos-api/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

var ownerSecretKey = config.GetEnvConfig("SECRET_KEY")

func OwnerAuthorizationRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SuccessHandler: ownerSuccessHandler,
		ErrorHandler:   ownerErrorHandler,
		SigningKey:     []byte(ownerSecretKey),
		SigningMethod:  "HS256",
	})
}

func ownerSuccessHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID, ok := claims["user_id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
			"msg":   "Error token cannot used this services.",
		})
	}
	Username, ok := claims["username"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
			"msg":   "Error token cannot used this services.",
		})
	}
	PrefixID, ok := claims["prefix_id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
			"msg":   "Error token cannot used this services.",
		})
	}
	RoleID, ok := claims["role_id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
			"msg":   "Error token cannot used this services.",
		})
	}
	if RoleID != "664b1d88474fd0eb97d02dff" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
			"msg":   "Error token cannot used this services.",
		})
	}
	c.Locals("user_id", ID)
	c.Locals("username", Username)
	c.Locals("prefix_id", PrefixID)
	c.Locals("role_id", RoleID)
	return c.Next()
}

func ownerErrorHandler(c *fiber.Ctx, e error) error {
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
		"msg":   e.Error(),
	})
	return nil
}
