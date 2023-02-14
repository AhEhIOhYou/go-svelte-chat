package application

import (
	"project-eighteen/pkg/server/domain/entities"
	"project-eighteen/pkg/server/domain/repository"
)

type ContactApp struct {
	contactRep repository.ContactRepository
}

var _ ContactAppInterface = &ContactApp{}

type ContactAppInterface interface {
	CreateContact(*entities.Contact) (*entities.Contact, error)
	GetLisOfContactsByUserID(string) ([]*entities.Contact, error)
	DeleteContact(string) error
	DeleteContactsByUserID(string) error
}

func (c *ContactApp) CreateContact(contact *entities.Contact) (*entities.Contact, error) {
	return c.contactRep.CreateContact(contact)
}

func (c *ContactApp) GetLisOfContactsByUserID(userID string) ([]*entities.Contact, error) {
	return c.contactRep.GetLisOfContactsByUserID(userID)
}

func (c *ContactApp) DeleteContact(contactID string) error {
	return c.contactRep.DeleteContact(contactID)
}

func (c *ContactApp) DeleteContactsByUserID(userID string) error {
	return c.contactRep.DeleteContactsByUserID(userID)
}
