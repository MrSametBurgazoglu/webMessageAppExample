package router

import (
	"github.com/gin-gonic/gin"
	"webMessageApp/controller"
)

func SetMessageRouter(engine *gin.Engine) {
	messageApi := engine.Group("/message")
	{
		messageApi.POST("/messages", controller.CreateMessage)
		messageApi.POST("/create_chat/", controller.CreateChat)
		messageApi.GET("/messages/:id", controller.FindMessage)
		messageApi.GET("/chat_messages/:id", controller.GetChatMessages)
	}
}
