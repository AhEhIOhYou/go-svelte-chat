package repository

import (
	"project-eighteen/pkg/server/domain/entities"
)

type ContactRepository interface {
	CreateContact(*entities.Contact) (*entities.Contact, error)
	GetListOfContactsByUserID(string) ([]*entities.Contact, error)
	DeleteContact(string) error
	DeleteContactsByUserID(string) error
}
