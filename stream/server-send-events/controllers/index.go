package ServerSendEventControllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	ServerSendEventsServices "pos-api/stream/server-send-events/services"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

var Validator = validator.New()

type EventInfo struct {
	Channel string `json:"channel" validate:"required"`
	Status  string `json:"status" validate:"required"`
	Message string `json:"message" validate:"required"`
}

func PublishNotification(c *fiber.Ctx) error {
	body := new(EventInfo)
	c.BodyParser(&body)
	err := Validator.Struct(body)
	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Field()+" is "+err.Tag())
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "validate error",
			"message": errors,
		})
	}
	json_data, _ := json.Marshal(body)
	ServerSendEventsServices.Broadcast(string(json_data))
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status":  "success",
		"message": "Message is Public.",
	})
}

func Notification(c *fiber.Ctx) error {
	prefix_id := c.Params("prefix_id")
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")
	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		eventChan := make(chan string)
		ServerSendEventsServices.Clients[eventChan] = struct{}{}

		for {
			data := <-eventChan
			evnet_info := EventInfo{}
			json.Unmarshal([]byte(data), &evnet_info)
			if evnet_info.Channel == prefix_id {
				type resp struct {
					Status  string `json:"status"`
					Message string `json:"message"`
				}
				json_data, _ := json.Marshal(resp{
					Status:  evnet_info.Status,
					Message: evnet_info.Message,
				})
				fmt.Fprintf(w, "data: notification: %s\n\n", json_data)
				err := w.Flush()
				if err != nil {
					fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)
					break
				}
			}
			time.Sleep(2 * time.Second)
		}
	}))
	return nil
}
