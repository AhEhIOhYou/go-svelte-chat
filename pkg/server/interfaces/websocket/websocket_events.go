package websocket

import (
	"log"
	"project-eighteen/pkg/server/constants"
	"project-eighteen/pkg/server/domain/entities"
)

func userOnline(c *ClientType, socketEvent entities.SocketEventType) {
	userID := socketEvent.Payload.(string)
	userDetails, err := c.hub.userApp.GetUserByID(userID)

	if (userDetails == &entities.User{} || err != nil) {
		log.Println("Invalid user with id: ", userID, " tried to join")
	} else {
		log.Println("User with id: ", userID, " online")

		// 1. Обновить статус пользователя в БД
		c.hub.userApp.UpdateUserOnlineStatus(userID, constants.UserOnline)

		// 2. Создать событие о том, что пользователь вошел в чат
		userOnlineEvent := entities.SocketEventType{
			Name: "user-update",
			Payload: map[string]interface{}{
				"type": "connected",
				"user": map[string]interface{}{
					"userID":   userDetails.ID,
					"username": userDetails.Username,
					"online":   constants.UserOnline,
				},
			},
		}

		// 3. Отправить событие о том, что пользователь вошел в чат всем контактам пользователя
		BroadcastToContacts(c.hub, userOnlineEvent, userID)
	}
}

func userDisconnect(c *ClientType, socketEvent entities.SocketEventType) {
	userID := socketEvent.Payload.(string)
	userDetails, err := c.hub.userApp.GetUserByID(userID)
	
	if (userDetails == &entities.User{} || err != nil) {
		log.Println("Invalid user with id: ", userID, " tried to disconnect")
	} else {
		log.Println("User with id: ", userID, " disconnected")

		// 1. Обновить статус пользователя в БД
		c.hub.userApp.UpdateUserOnlineStatus(userID, constants.UserOffline)

		// 2. Создать событие о том, что пользователь вышел из чата
		userOfflineEvent := entities.SocketEventType{
			Name: "user-update",
			Payload: map[string]interface{}{
				"type": "disconnected",
				"user": map[string]interface{}{
					"userID":   userDetails.ID,
					"username": userDetails.Username,
				},
			},
		}

		// 3. Отправить событие о том, что пользователь вышел из чата всем контактам пользователя
		BroadcastToContacts(c.hub, userOfflineEvent, userID)
	}
}

func messageStore(c *ClientType, socketEvent entities.SocketEventType) {
	message := socketEvent.Payload.(entities.Message)

	if (message == entities.Message{}) {
		log.Println("Invalid message")
	} else {
		// 1. Сохранить сообщение в БД
		messageID, err := c.hub.messageApp.StoreMessage(&message)
		if err != nil {
			log.Println("Failed to store message: ", err)
			return
		}
		log.Println("Message from user with id: ", message.FromID, " to chat with id: ", message.ChatID, " stored")

		// 2. Создать событие о том, что пользователь отправил сообщение
		messageEvent := entities.SocketEventType{
			Name: "message",
			Payload: map[string]interface{}{
				"messageID": messageID,
				"fromID":  message.FromID,
				"fromName": message.FromName,
				"chatID":  message.ChatID,
				"chatName": message.ChatName,
				"message": message.Message,
				"createdAt": message.CreatedAt,
			},
		}

		// 3. Отправить событие о том, что пользователь отправил сообщение всем участникам чата
		BroadcastToChat(c.hub, messageEvent, message.ChatID)
	}
}