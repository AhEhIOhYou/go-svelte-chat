package entities

type SocketEvent struct {
	EventName    string      `json:"eventName"`
	EventPayload interface{} `json:"eventPayload"`
}