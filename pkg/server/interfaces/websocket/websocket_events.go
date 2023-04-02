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
	// 0. Преобразовать событие в сообщение
	messageRaw := socketEvent.Payload.(map[string]interface{})

	message := entities.Message{
		FromID:    messageRaw["fromUserID"].(string),
		FromName:  messageRaw["fromUserName"].(string),
		ChatID:    messageRaw["chatID"].(string),
		Message:   messageRaw["message"].(string),
		CreatedAt: messageRaw["createdAt"].(string),
	}
	authorID := message.FromID

	if (message == entities.Message{}) {
		log.Println("Invalid message")
	} else {
		// 1. Сохранить сообщение в БД
		message, err := c.hub.messageApp.StoreMessage(&message)
		if err != nil {
			log.Println("Failed to store message: ", err)
			return
		}
		log.Println("Message from user with id: ", authorID, " to chat with id: ", message.ChatID, " stored")

		// 2. Создать событие о том, что пользователь отправил сообщение
		messageEvent := entities.SocketEventType{
			Name:    "message",
			Payload: message,
		}

		// 3. Отправить событие о том, что пользователь отправил сообщение всем участникам чата
		BroadcastToChat(c.hub, messageEvent, message.ChatID)
	}
}

func notifyNewChat(c *ClientType, socketEvent entities.SocketEventType) {
	// 0. Преобразовать событие в сообщение
	chatRaw := socketEvent.Payload.(map[string]interface{})
	var participants []string

	for _,item:= range chatRaw["participants"].([]interface{}) {
		participants = append(participants, item.(string))
	}

	chat := entities.Chat{
		ParticipantsID: participants,
		Name:           chatRaw["name"].(string),
	}

	// 1. Сохранить чат в БД
	newChat, err := c.hub.chatApp.CreateChat(&chat)
	if err != nil {
		log.Println("Failed to create chat: ", err)
		return
	}
	log.Println("New chat created")

	// 2. Создать событие о том, что пользователь создал чат
	chatEvent := entities.SocketEventType{
		Name:    "new-chat",
		Payload: newChat,
	}

	// 3. Отправить событие о том, что пользователь создал чат всем участникам чата
	BroadcastToChat(c.hub, chatEvent, newChat.ID)
}
