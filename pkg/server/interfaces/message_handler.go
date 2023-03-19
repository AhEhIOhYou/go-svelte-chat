package interfaces

import (
	"fmt"
	"net/http"
	"project-eighteen/pkg/constants"
	"project-eighteen/pkg/server/application"
	"project-eighteen/pkg/server/domain/entities"
	"time"

	"github.com/gin-gonic/gin"
)

type Message struct {
	messageApp application.MessageAppInterface
	chatApp    application.ChatAppInterface
}

func NewMessageHandler(messageApp application.MessageAppInterface, chatApp application.ChatAppInterface) *Message {
	return &Message{
		messageApp: messageApp,
		chatApp:    chatApp,
	}
}

func (m *Message) StoreNewMessage(ctx *gin.Context) {
	var message entities.Message

	err := ctx.BindJSON(&message)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	message.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	validateErr := message.Validate()
	if validateErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, validateErr),
		})
		return
	}

	chat, err := m.chatApp.GetChatByID(message.ChatID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	if chat.ID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.ChatDoesNotExist,
		})
		return
	}

	newMessage, err := m.messageApp.StoreMessage(&message)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":     constants.Successful,
		"new-message": newMessage,
	})
}

func (m *Message) DeleteMessage(ctx *gin.Context) {
	messageId := ctx.Param("messageId")

	message, err := m.messageApp.GetMessageByID(messageId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	createdAt, err := time.Parse("2006-01-02 15:04:05", message.CreatedAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	if time.Since(createdAt).Hours() > 12 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.MessageCannotBeDeleted,
		})
		return
	}

	err = m.messageApp.DeleteMessage(messageId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": constants.Successful,
	})
}

func (m *Message) EditMessage(ctx *gin.Context) {
	var message entities.Message

	err := ctx.BindJSON(&message)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	validateErr := message.Validate()
	if validateErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, validateErr),
		})
		return
	}

	createdAt, err := time.Parse("2006-01-02 15:04:05", message.CreatedAt)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	if time.Since(createdAt).Hours() > 12 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.MessageCannotBeEdited,
		})
		return
	}

	editedMessage, err := m.messageApp.UpdateMessage(&message)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":        constants.Successful,
		"edited-message": editedMessage,
	})
}

func (m *Message) DeleteMessagesByChatID(ctx *gin.Context) {
	chatId := ctx.Param("chatId")

	err := m.messageApp.DeleteMessagesByChatID(chatId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": constants.Successful,
	})
}
