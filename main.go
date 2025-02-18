package main

import (
	"chat-tool-calling/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	ctrl := controllers.NewChatController()

	router.POST("/api/chat", ctrl.GetChat)

	router.Run(":8080")
}
