package server

import (
	"log"

	"project-eighteen/pkg/database"
	"project-eighteen/pkg/server/httpserver"
	"project-eighteen/pkg/server/httpserver/middleware"
	"project-eighteen/pkg/server/websocket"

	"github.com/gin-gonic/gin"
)

func Start() {

	database.DBConnect()

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	hub := websocket.NewHub()
	go hub.Run()

	router.GET("/ws", hub.ServeWs)
	log.Println("Server websocket started")

	router.GET("/", httpserver.Home)
	router.GET("/is-username-available/:username", httpserver.IsUsernameAvailable)
	router.POST("/registration", httpserver.Registration)
	router.POST("/login", httpserver.Login)
	router.GET("/user-session-check", httpserver.UserSessionCheck)
	router.GET("/chat/:userFromId/:userToId", httpserver.GetMessagesHandler)
	router.GET("/search-user", httpserver.SearchUser)
	router.POST("/add-contact", httpserver.AddContact)
	router.GET("/contacts", httpserver.GetContacts)

	log.Println("Server started")

	log.Fatal(router.Run(":8080"))
}
