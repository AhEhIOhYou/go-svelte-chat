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
	/*
		Created chat with participants and return chat entity. If chat already exist, return chat entity.
		If two practitioners in chat, create with type ChatTypeDialog else ChatTypeGroup.
	*/
	CreateChat(*entities.Chat) (*entities.Chat, error)
	// Check exist dialog with participants and return it if exist.
	CheckExsistDialog([]string) (*entities.Chat, error)
	// Return chat by id.
	GetChatByID(string) (*entities.Chat, error)
	// Return all chats with user id.
	GetChatsByParticipantID(string) ([]*entities.Chat, error)
	// Update chat.
	UpdateChat(*entities.Chat) (*entities.Chat, error)
	// Add participant to group chat.
	AddParticipantToChat(string, string) error
	// Delete participant from group chat.
	DeleteParticipantFromChat(string, string) error
	// Delete chat.
	DeleteChat(string) error
}

func (c *ChatApp) CreateChat(chat *entities.Chat) (*entities.Chat, error) {
	return c.chatRep.CreateChat(chat)
}

func (c *ChatApp) CheckExsistDialog(participantsID []string) (*entities.Chat, error) {
	return c.chatRep.CheckExsistDialog(participantsID)
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
