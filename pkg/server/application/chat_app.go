package application

import (
	"project-eighteen/pkg/server/domain/entities"
	"project-eighteen/pkg/server/domain/repository"
)

type ChatApp struct {
	chatRep repository.ChatRepository
}

var _ ChatAppInterface = &ChatApp{}

type ChatAppInterface interface {
	CreateChat(*entities.Chat) (*entities.Chat, error)
	GetChatByID(string) (*entities.Chat, error)
	GetChatsByParticipantID(string) ([]*entities.Chat, error)
	UpdateChat(*entities.Chat) (*entities.Chat, error)
	AddParticipantToChat(string, string) error
	DeleteParticipantFromChat(string, string) error
	DeleteChat(string) error
}

func (c *ChatApp) CreateChat(chat *entities.Chat) (*entities.Chat, error) {
	return c.chatRep.CreateChat(chat)
}

func (c *ChatApp) GetChatByID(chatID string) (*entities.Chat, error) {
	return c.chatRep.GetChatByID(chatID)
}

func (c *ChatApp) GetChatsByParticipantID(userID string) ([]*entities.Chat, error) {
	return c.chatRep.GetChatsByParticipantID(userID)
}

func (c *ChatApp) UpdateChat(chat *entities.Chat) (*entities.Chat, error) {
	return c.chatRep.UpdateChat(chat)
}

func (c *ChatApp) AddParticipantToChat(chatID string, userID string) error {
	return c.chatRep.AddParticipantToChat(chatID, userID)
}

func (c *ChatApp) DeleteParticipantFromChat(chatID string, userID string) error {
	return c.chatRep.DeleteParticipantFromChat(chatID, userID)
}

func (c *ChatApp) DeleteChat(chatID string) error {
	return c.chatRep.DeleteChat(chatID)
}