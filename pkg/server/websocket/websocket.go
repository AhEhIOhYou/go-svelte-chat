package websocket

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"project-eighteen/pkg/server/httpserver"
	"project-eighteen/pkg/server/structs"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type ClientType struct {
	hub    *HubType
	wsConn *websocket.Conn
	send   chan structs.SocketEventType
	userID string
}

func unRegisterAndCloseConn(c *ClientType) {
	c.hub.unregister <- c
	c.wsConn.Close()
}

func setSocketPayloadReadConfig(c *ClientType) {
	c.wsConn.SetReadLimit(maxMessageSize)
	c.wsConn.SetReadDeadline(time.Now().Add(pongWait))
	c.wsConn.SetPongHandler(func(string) error { c.wsConn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
}

func handleSocketPayloadEvents(c *ClientType, socketEvent structs.SocketEventType) {
	type chatlistResponseType struct {
		Type     string      `json:"type"`
		Chatlist interface{} `json:"chatlist"`
	}

	switch socketEvent.Name {
	case "join":
		userID := socketEvent.Payload.(string)
		userDetails := httpserver.GetUserByUserID(userID)
		if userDetails == (structs.UserDetailsType{}) {
			log.Println("Invalid user with id: ", userID, " tried to join")
		} else {
			if userDetails.Online == "N" {
				log.Println("User with id: ", userID, " joined")
			} else {
				newUserOnlinePayload := structs.SocketEventType{
					Name: "chatlist-res",
					Payload: chatlistResponseType{
						Type: "new-user-joined",
						Chatlist: structs.UserDetailsResponsePayloadType{
							UserID:   userDetails.ID,
							Username: userDetails.Username,
							Online:   userDetails.Online,
						},
					},
				}

				BroadcastToAllExceptMe(c.hub, newUserOnlinePayload, userDetails.ID)

				allOnlineUsersPayload := structs.SocketEventType{
					Name: "chatlist-res",
					Payload: chatlistResponseType{
						Type:     "my-chatlist",
						Chatlist: httpserver.GetAllOnlineUsers(userDetails.ID),
					},
				}

				SendToClient(c.hub, allOnlineUsersPayload, userDetails.ID)
			}
		}
	case "disconnect":
		if socketEvent.Payload != nil {
			userID := socketEvent.Payload.(string)
			userDetails := httpserver.GetUserByUserID(userID)
			httpserver.UpdateUserOnlineStatus(userID, "N")

			BroadcastToAll(c.hub, structs.SocketEventType{
				Name: "chatlist-res",
				Payload: chatlistResponseType{
					Type: "user-disconnected",
					Chatlist: structs.UserDetailsResponsePayloadType{
						UserID:   userDetails.ID,
						Username: userDetails.Username,
						Online:   "N",
					},
				},
			})
		}
	case "message":
		message := socketEvent.Payload.(map[string]interface{})["message"].(string)
		fromUserID := socketEvent.Payload.(map[string]interface{})["fromUserID"].(string)
		toUserID := socketEvent.Payload.(map[string]interface{})["toUserID"].(string)

		if message != "" && fromUserID != "" && toUserID != "" {
			messagePayload := structs.MessagePayloadType{
				From:    fromUserID,
				To:      toUserID,
				Message: message,
			}
			httpserver.StoreNewChatMessage(messagePayload)

			allOnlineUsersPayload := structs.SocketEventType{
				Name:    "message-res",
				Payload: messagePayload,
			}
			SendToClient(c.hub, allOnlineUsersPayload, toUserID)
		}
	}
}

func (c *ClientType) readPump() {
	var socketEvent structs.SocketEventType
	defer unRegisterAndCloseConn(c)
	setSocketPayloadReadConfig(c)

	for {
		_, payload, err := c.wsConn.ReadMessage()

		decoder := json.NewDecoder(bytes.NewReader(payload))
		decoderErr := decoder.Decode(&socketEvent)

		if decoderErr != nil {
			log.Println("Error decoding socket event: ", decoderErr)
			break
		}

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error: %v", err)
			}
			break
		}

		handleSocketPayloadEvents(c, socketEvent)
	}
}

func (c *ClientType) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.wsConn.Close()
	}()

	for {
		select {
		case payload, ok := <-c.send:
			reqBodyBytes := new(bytes.Buffer)
			json.NewEncoder(reqBodyBytes).Encode(payload)
			finalPayload := reqBodyBytes.Bytes()

			c.wsConn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.wsConn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.wsConn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(finalPayload)

			n := len(c.send)
			for i := 0; i < n; i++ {
				json.NewEncoder(reqBodyBytes).Encode(<-c.send)
				w.Write(reqBodyBytes.Bytes())
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.wsConn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.wsConn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func CreateNewSocketUser(hub *HubType, conn *websocket.Conn, userID string) {
	client := &ClientType{
		hub:    hub,
		wsConn: conn,
		send:   make(chan structs.SocketEventType),
		userID: userID,
	}

	go client.writePump()
	go client.readPump()

	client.hub.register <- client
}

func HandleUserRegisterEvent(hub *HubType, client *ClientType) {
	hub.clients[client] = true
	handleSocketPayloadEvents(client, structs.SocketEventType{
		Name:    "join",
		Payload: client.userID,
	})
}

func HandleUserDisconnectEvent(hub *HubType, client *ClientType) {
	if _, ok := hub.clients[client]; ok {
		handleSocketPayloadEvents(client, structs.SocketEventType{
			Name:    "disconnect",
			Payload: client.userID,
		})
		delete(hub.clients, client)
		close(client.send)
	}
}

func SendToClient(hub *HubType, payload structs.SocketEventType, userID string) {
	for client := range hub.clients {
		if client.userID == userID {
			select {
			case client.send <- payload:
			default:
				close(client.send)
				delete(hub.clients, client)
			}
		}
	}
}

func BroadcastToAll(hub *HubType, payload structs.SocketEventType) {
	for client := range hub.clients {
		select {
		case client.send <- payload:
		default:
			close(client.send)
			delete(hub.clients, client)
		}
	}
}

func BroadcastToAllExceptMe(hub *HubType, payload structs.SocketEventType, userID string) {
	for client := range hub.clients {
		if client.userID != userID {
			select {
			case client.send <- payload:
			default:
				close(client.send)
				delete(hub.clients, client)
			}
		}
	}
}
