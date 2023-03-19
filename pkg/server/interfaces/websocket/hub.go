package websocket

import (
	"log"
	"net/http"
	"project-eighteen/pkg/server/application"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type HubType struct {
	clients    map[*ClientType]bool
	register   chan *ClientType
	unregister chan *ClientType
	userApp    application.UserAppInterface
	chatApp    application.ChatAppInterface
	contactApp application.ContactApp
	messageApp application.MessageApp
}

func NewHub(userApp application.UserAppInterface, chatApp application.ChatAppInterface, contactApp application.ContactApp, messageApp application.MessageApp) *HubType {
	return &HubType{
		register:   make(chan *ClientType),
		unregister: make(chan *ClientType),
		clients:    make(map[*ClientType]bool),
		userApp:    userApp,
		chatApp:    chatApp,
		contactApp: contactApp,
		messageApp: messageApp,
	}
}

func (h *HubType) Run() {
	for {
		select {
		case client := <-h.register:
			HandleUserRegisterEvent(h, client)
		case client := <-h.unregister:
			HandleUserDisconnectEvent(h, client)
		}
	}
}

func (h *HubType) ServeWs(ctx *gin.Context) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	userId := ctx.Query("userID")

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	wsConn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("Error upgrading to websocket: ", err)
		return
	}

	CreateNewSocketUser(h, wsConn, userId)
}
