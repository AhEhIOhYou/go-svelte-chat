package server

import (
	"log"

	"project-eighteen/pkg/server/httpserver"
	"project-eighteen/pkg/server/httpserver/middleware"
	"project-eighteen/pkg/server/websocket"

	"github.com/gin-gonic/gin"
)

func StartHttpServer() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	router.GET("/", httpserver.Home)
	router.GET("/is-username-available/:username", httpserver.IsUsernameAvailable)
	router.POST("/registration", httpserver.Registration)
	router.POST("/login", httpserver.Login)
	router.GET("user-session-check", httpserver.UserSessionCheck)
	router.GET("chat/:userFromId/:userToId", httpserver.GetMessagesHandler)

	log.Println("Server started on port 8080")

	log.Fatal(router.Run(":8080"))
}

func StartWebSocketServer() {
	hub := websocket.NewHub()
	go hub.Run()

	router := gin.Default()
	router.GET("/ws", hub.ServeWs)

	log.Println("Server websocket started on port 8081")

	log.Fatal(router.Run(":8081"))
}
