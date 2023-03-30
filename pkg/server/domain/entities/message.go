package entities

import (
	"errors"
	"project-eighteen/pkg/server/constants"
)

type Message struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	ChatID    string `json:"chatID" bson:"chat_id,omitempty"`
	FromID    string `json:"fromUserID" bson:"from_id,omitempty"`
	FromName  string `json:"fromUserName" bson:"from_name,omitempty"`
	Message   string `json:"message" bson:"message,omitempty"`
	CreatedAt string `json:"createdAt" bson:"created_at,omitempty"`
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
