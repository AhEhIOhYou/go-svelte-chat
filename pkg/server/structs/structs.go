package structs

type UserDetailsType struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string
	Password string
	Online   string
	SocketID string
}

type ChatType struct {
	ID      string `json:"id" bson:"_id,omitempty"`
	Message string `json:"message"`
	From    string `json:"fromUserID"`
	To      string `json:"toUserID"`
}

type UserDetailsRequestPayloadType struct {
	Username string
	Password string
}

type UserDetailsResponsePayloadType struct {
	Username string `json:"username"`
	UserID   string `json:"userId"`
	Online   string `json:"online"`
}

type SocketEventType struct {
	EventName    string      `json:"eventName"`
	EventPayload interface{} `json:"eventPayload"`
}

type MessagePayloadType struct {
	From    string `json:"fromUserID"`
	To      string `json:"toUserID"`
	Message string `json:"message"`
}
