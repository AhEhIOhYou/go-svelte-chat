package entities

import (
	"errors"
	"project-eighteen/pkg/server/constants"
)

type Chat struct {
	ID             string   `json:"id" bson:"_id,omitempty"`
	Type           int      `json:"type" bson:"type"`
	ParticipantsID []string `json:"participants" bson:"participants"`
	Name           string   `json:"name" bson:"name"`
}

type Chats []Chat

func (chat *Chat) Validate() error {
	if chat.Name == "" {
		return errors.New(constants.ChatNameCantBeEmpty)
	}
	if len(chat.ParticipantsID) < 2 {
		return errors.New(constants.ChatMustHaveAtLeastTwoParticipants)
	}

	return nil
}
