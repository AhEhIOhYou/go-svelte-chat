package entities

import (
	"errors"
	"project-eighteen/pkg/constants"
)

type Contact struct {
	ID string `json:"id" bson:"_id,omitempty"`
	UserID string `json:"userID"`
	ContactID string `json:"contactUserID"`
}

type ContactRequest struct {
	UserID string `json:"userID"`
	ContactID string `json:"contactUserID"`
}

type ContactResponse struct {
	UserID string `json:"userID"`
	ContactID string `json:"contactUserID"`
	ContactUsername string `json:"contactUsername"`
	Online string `json:"online"`
}

func (contact *Contact) Validate() error {
	if contact.UserID == "" {
		return errors.New(constants.UserIDCantBeEmpty)
	}
	if contact.ContactID == "" {
		return errors.New(constants.ContactIDCantBeEmpty)
	}

	return nil
}