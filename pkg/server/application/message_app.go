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
	// Store message in database and return it
	StoreMessage(*entities.Message) (*entities.Message, error)
	// Get messages by chat id. Limit and offset are used for pagination.
	GetMessagesByChatID(string, int64, int64) ([]*entities.Message, error)
	// Get message by id
	GetMessageByID(string) (*entities.Message, error)
	// Update message in database and return it
	UpdateMessage(*entities.Message) (*entities.Message, error)
	// Delete message by id
	DeleteMessage(string) error
	// Delete all messages by chat id
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
