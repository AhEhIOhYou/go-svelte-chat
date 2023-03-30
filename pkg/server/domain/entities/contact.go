package entities

import (
	"errors"
	"project-eighteen/pkg/server/constants"
)

type Contact struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	UserID    string `json:"userID" bson:"user_id"`
	ContactUsername string `json:"contactUsername" bson:"contact_username"`
	ContactID string `json:"contactUserID" bson:"contact"`
}

func (contact *Contact) Validate() error {
	if contact.UserID == "" {
		return errors.New(constants.UserIDCantBeEmpty)
	}
	if contact.ContactUsername == "" {
		return errors.New(constants.ContactUsernameCantBeEmpty)
	}
	if contact.ContactID == "" {
		return errors.New(constants.ContactIDCantBeEmpty)
	}

	return nil
}
