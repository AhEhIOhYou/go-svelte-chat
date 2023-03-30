package server

import (
	"log"
	"os"
	"project-eighteen/pkg/server/infrastructure/persistence"
	"project-eighteen/pkg/server/interfaces"
	"project-eighteen/pkg/server/interfaces/middleware"
	"project-eighteen/pkg/server/interfaces/websocket"

	"github.com/gin-gonic/gin"
)

func Start() {

	uri := os.Getenv("MONGODB_URI")

	services, err := persistence.NewRepo(uri)
	if err != nil {
		log.Fatal(err)
	}

	users := interfaces.NewUsersHandler(services.UserRepository)
	chat := interfaces.NewChatHandler(services.ChatRepository, services.MessageRepository)
	contacts := interfaces.NewContactsHandler(services.ContactRepository, services.UserRepository)
	messages := interfaces.NewMessageHandler(services.MessageRepository, services.ChatRepository)
	index := interfaces.Ping

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	hub := websocket.NewHub(services.UserRepository, services.ChatRepository, services.ContactRepository, services.MessageRepository)
	go hub.Run()

	router.GET("/ws", hub.ServeWs)

	u := router.Group("api/users")
	{
		u.POST("/register", users.Register)
		u.POST("/login", users.Login)
		u.GET("/is-username-available", users.IsUsernameAvailable)
		u.GET("/search", users.SearchUser)
		u.GET("/logout", users.Logout)
		u.GET("/:user-id", users.GetByID)
	}

	ch := router.Group("api/chats")
	{
		ch.POST("/create", chat.CreateChat)
		ch.GET("/chats-by-user", chat.GetChatsByUser)
		ch.GET("/:chat-id", chat.GetChatData)
		ch.POST("/dialog", chat.GetDialog)
	}

	con := router.Group("api/contacts")
	{
		con.POST("/add", contacts.AddContact)
		con.GET("/:user-id", contacts.GetContacts)
		con.DELETE("/delete", contacts.DeleteContact)
		con.DELETE("/delete-all", contacts.ClearAllUserContacts)
	}

	mes := router.Group("api/messages")
	{
		mes.POST("/send", messages.StoreNewMessage)
		mes.DELETE("/delete", messages.DeleteMessage)
		mes.PUT("/edit", messages.EditMessage)
		mes.DELETE("/delete-all", messages.DeleteMessagesByChatID)
	}

	router.GET("/ping", index)

	router.RunTLS(":8081", os.Getenv("CRT"), os.Getenv("KEY"))
}