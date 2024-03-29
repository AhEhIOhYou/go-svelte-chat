package repository

import (
	"project-eighteen/pkg/server/domain/entities"
)

type ChatRepository interface {
	CreateChat(*entities.Chat) (*entities.Chat, error)
	CheckExsistDialog([]string) (*entities.Chat, error)
	GetChatByID(string) (*entities.Chat, error)
	GetChatsByParticipantID(string) ([]*entities.Chat, error)
	UpdateChat(*entities.Chat) (*entities.Chat, error)
	AddParticipantToChat(string, string) error
	DeleteParticipantFromChat(string, string) error
	DeleteChat(string) error
}