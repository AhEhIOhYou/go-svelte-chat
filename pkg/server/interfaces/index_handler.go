package interfaces

import "github.com/gin-gonic/gin"

type Index struct{}

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}