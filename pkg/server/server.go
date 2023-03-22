package server

import (
	"log"
	"os"

	"project-eighteen/pkg/server/httpserver/middleware"
	"project-eighteen/pkg/server/infrastructure/persistence"
	"project-eighteen/pkg/server/interfaces/websocket"

	"github.com/gin-gonic/gin"
)

func Start() {

	uri := os.Getenv("MONGODB_URI")

	services, err := persistence.NewRepo(uri)
	if err != nil {
		log.Fatal(err)
	}

	// users := interfaces.NewUsersHandler(services.UserRepository)
	// chat := interfaces.NewChatHandler(services.ChatRepository, services.MessageRepository)
	// contacts := interfaces.NewContactsHandler(services.ContactRepository)
	// messages := interfaces.NewMessageHandler(services.MessageRepository, services.ChatRepository)
	// index := interfaces.Ping

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	// websocket new hub
	hub := websocket.NewHub(services.UserRepository, services.ChatRepository, services.ContactRepository, services.MessageRepository)
	go hub.Run()

	// websocket routes
	router.GET("/ws", hub.ServeWs)

	router.Run(":8081")

}