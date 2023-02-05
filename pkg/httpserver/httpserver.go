package httpserver

import (
	"log"
	"project-eighteen/pkg/httpserver/middleware"
	"project-eighteen/pkg/redisrepo"

	"github.com/gin-gonic/gin"
)


func StartServer() {
	redisClient := redisrepo.InitialiseRedis()
	defer redisClient.Close()

	redisrepo.CreateFetchChatBetweenIndex()

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/register", registerHandler)
	router.POST("/login", loginHandler)
	router.GET("/verify/:username", verifyContactHandler)
	router.GET("/chat-history", chatHistoryHandler)
	router.GET("/contact-list", contactListHandler)

	log.Fatal(router.Run(":8080"))
}
