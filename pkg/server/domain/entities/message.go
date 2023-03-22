package entities

import (
	"errors"
	"project-eighteen/pkg/constants"
)

type Message struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	ChatID    string `json:"chatID"`
	ChatName  string `json:"chatName"`
	FromID      string `json:"fromUserID"`
	FromName  string `json:"fromUserName"`
	Message   string `json:"message"`
	CreatedAt string `json:"createdAt"`
}

type MessageRequest struct {
	ChatID  string `json:"chatID"`
	From    string `json:"fromUserID"`
	Message string `json:"message"`
}

func (message *Message) Validate() error {
	if message.Message == "" {
		return errors.New(constants.MessageCantBeEmpty)
	}
	if message.ChatID == "" {
		return errors.New(constants.ChatIDCantBeEmpty)
	}
	if message.FromID == "" {
		return errors.New(constants.FromUserIDCantBeEmpty)
	}

	return nil
}
