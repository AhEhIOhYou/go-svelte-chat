package entities

import (
	"errors"
	"project-eighteen/pkg/constants"
)

type Chat struct {
	ID string `json:"id" bson:"_id,omitempty"`
	ParticipantsID []string `json:"participantsID"`
	Name string `json:"name"`
}

type ChatRequest struct {
	ParticipantsID []string `json:"participantsID"`
	Name string `json:"name"`
}

func (chat *Chat) Validate() error {
	if chat.Name == "" {
		return errors.New(constants.ChatNameCantBeEmpty)
	}
	if len(chat.ParticipantsID) < 2 {
		return errors.New(constants.ChatMustHaveAtLeastTwoParticipants)
	}

	return nil
}