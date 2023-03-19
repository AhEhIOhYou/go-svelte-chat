package repository

import (
	"project-eighteen/pkg/server/domain/entities"
)

type MessageRepository interface {
	StoreMessage(*entities.Message) (*entities.Message, error)
	GetMessagesByChatID(string, int64, int64) ([]*entities.Message, error)
	GetMessageByID(string) (*entities.Message, error)
	UpdateMessage(*entities.Message) (*entities.Message, error)
	DeleteMessage(string) error
	DeleteMessagesByChatID(string) error
}