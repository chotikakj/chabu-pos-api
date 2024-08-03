package ServerSendEventRouter

import (
	Controllers "pos-api/stream/server-send-events/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	app := router.Group("sse")
	app.Get("/notification/:prefix_id", Controllers.Notification)
	app.Post("/publish-notification", Controllers.PublishNotification)
}
