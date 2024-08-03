package ServerSendEventsServices

import "encoding/json"

var Clients = make(map[chan string]struct{})

func Broadcast(data string) {
	for client := range Clients {
		client <- data
	}
}

type EventInfo struct {
	Channel string `json:"channel" validate:"required"`
	Status  string `json:"status" validate:"required"`
	Message string `json:"message" validate:"required"`
}

func Notification(info EventInfo) {
	json_data, _ := json.Marshal(info)
	Broadcast(string(json_data))
}
