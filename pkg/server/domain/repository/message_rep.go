package repository

import (
	"project-eighteen/pkg/server/domain/entities"
)

type MessageRepository interface {
	CreateMessage(*entities.Message) (*entities.Message, error)
	GetMessageByChatID(string, int, int) ([]*entities.Message, error)
	UpdateMessage(*entities.Message) (*entities.Message, error)
	DeleteMessage(string) error
	DeleteMessagesByChatID(string) error
}