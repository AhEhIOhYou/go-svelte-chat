package application

import (
	"project-eighteen/pkg/server/domain/entities"
	"project-eighteen/pkg/server/domain/repository"
)

type MessageApp struct {
	messageRepository repository.MessageRepository
}

var _ MessageAppInterface = &MessageApp{}

type MessageAppInterface interface {
	CreateMessage(*entities.Message) (*entities.Message, error)
	GetMessageByChatID(string, int, int) ([]*entities.Message, error)
	UpdateMessage(*entities.Message) (*entities.Message, error)
	DeleteMessage(string) error
	DeleteMessagesByChatID(string) error
}

func (m *MessageApp) CreateMessage(message *entities.Message) (*entities.Message, error) {
	return m.messageRepository.CreateMessage(message)
}

func (m *MessageApp) GetMessageByChatID(chatID string, limit int, offset int) ([]*entities.Message, error) {
	return m.messageRepository.GetMessageByChatID(chatID, limit, offset)
}

func (m *MessageApp) UpdateMessage(message *entities.Message) (*entities.Message, error) {
	return m.messageRepository.UpdateMessage(message)
}

func (m *MessageApp) DeleteMessage(messageID string) error {
	return m.messageRepository.DeleteMessage(messageID)
}

func (m *MessageApp) DeleteMessagesByChatID(chatID string) error {
	return m.messageRepository.DeleteMessagesByChatID(chatID)
}
