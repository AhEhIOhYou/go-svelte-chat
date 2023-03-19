package entities

type SocketEventType struct {
	Message string      `json:"message"`
	Name    string      `json:"name"`
	Payload interface{} `json:"payload"`
}
