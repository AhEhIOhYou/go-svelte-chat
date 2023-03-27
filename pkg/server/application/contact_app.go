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
	IsContact(string, string) (bool, error)
	GetListOfContactsByUserID(string) ([]*entities.Contact, error)
	DeleteContact(string, string) error
	DeleteContactsByUserID(string) error
}

func (c *ContactApp) CreateContact(contact *entities.Contact) (*entities.Contact, error) {
	return c.contactRep.CreateContact(contact)
}

func (c *ContactApp) IsContact(userID, contactID string) (bool, error) {
	return c.contactRep.IsContact(userID, contactID)
}

func (c *ContactApp) GetListOfContactsByUserID(userID string) ([]*entities.Contact, error) {
	return c.contactRep.GetListOfContactsByUserID(userID)
}

func (c *ContactApp) DeleteContact(userID, contactID string) error {
	return c.contactRep.DeleteContact(userID, contactID)
}

func (c *ContactApp) DeleteContactsByUserID(userID string) error {
	return c.contactRep.DeleteContactsByUserID(userID)
}
