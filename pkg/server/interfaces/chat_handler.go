package interfaces

import (
	"fmt"
	"net/http"
	"project-eighteen/pkg/constants"
	"project-eighteen/pkg/server/application"
	"project-eighteen/pkg/server/domain/entities"
	"project-eighteen/pkg/server/infrastructure/utils"

	"github.com/gin-gonic/gin"
)

type Chat struct {
	chatApp    application.ChatAppInterface
	messageApp application.MessageAppInterface
}

func NewChatHandler(chatApp application.ChatAppInterface, messageApp application.MessageAppInterface) *Chat {
	return &Chat{
		chatApp:    chatApp,
		messageApp: messageApp,
	}
}

func (c *Chat) CreateChat(ctx *gin.Context) {
	var chat entities.Chat

	err := ctx.BindJSON(&chat)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	validateErr := chat.Validate()
	if validateErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	if len(chat.ParticipantsID) == 2 {
		chatExists, err := c.chatApp.CheckExsistDialog(chat.ParticipantsID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf(constants.Failed, err),
			})
			return
		}

		if chatExists.ID != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"message": constants.ChatAlreadyExists,
				"chat":    chatExists,
			})
			return
		}
	}

	newChat, err := c.chatApp.CreateChat(&chat)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": constants.Successful,
		"chat":    newChat,
	})
}

func (c *Chat) GetChatsByUser(ctx gin.Context) {
	userID := ctx.Param("userID")
	if !utils.CheckAlphaNumeric(userID) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.UserIDInvalid,
		})
	}

	chats, err := c.chatApp.GetChatsByParticipantID(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": constants.Successful,
		"chats":   chats,
	})
}

func (c *Chat) GetChatData(ctx gin.Context) {
	chatId := ctx.Param("chatID")
	if !utils.CheckAlphaNumeric(chatId) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.ChatDoesNotExist,
		})
	}

	chat, err := c.chatApp.GetChatByID(chatId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	messages, err := c.messageApp.GetMessagesByChatID(chatId, 0, 100)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(constants.Failed, err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  constants.Successful,
		"chat":     chat,
		"messages": messages,
	})
}