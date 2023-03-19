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
	StoreMessage(*entities.Message) (*entities.Message, error)
	GetMessagesByChatID(string, int64, int64) ([]*entities.Message, error)
	GetMessageByID(string) (*entities.Message, error)
	UpdateMessage(*entities.Message) (*entities.Message, error)
	DeleteMessage(string) error
	DeleteMessagesByChatID(string) error
}

func (m *MessageApp) StoreMessage(message *entities.Message) (*entities.Message, error) {
	return m.messageRepository.StoreMessage(message)
}

func (m *MessageApp) GetMessagesByChatID(chatID string, limit int64, offset int64) ([]*entities.Message, error) {
	return m.messageRepository.GetMessagesByChatID(chatID, limit, offset)
}

func (m *MessageApp) GetMessageByID(messageID string) (*entities.Message, error) {
	return m.messageRepository.GetMessageByID(messageID)
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
