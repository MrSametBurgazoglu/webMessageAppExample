package router

import (
	"github.com/gin-gonic/gin"
	"webMessageApp/controller"
)

func SetMessageRouter(engine *gin.Engine) {
	groupApi := engine.Group("/message")
	{
		groupApi.GET("/messages", controller.FindMessages)
		groupApi.POST("/messages", controller.CreateMessage)
		groupApi.GET("/messages/:id", controller.FindMessage)
		groupApi.GET("/chat_messages/:id", controller.GetChatMessages)
		groupApi.PATCH("/messages/:id", controller.UpdateMessage)
		groupApi.DELETE("/messages/:id", controller.DeleteMessage)
	}
}
