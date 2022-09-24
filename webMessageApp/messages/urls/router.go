package urls

import (
	"github.com/gin-gonic/gin"
	"webMessageApp/messages/controller"
)

func SetRouter(engine *gin.Engine) {
	groupApi := engine.Group("/message")
	{
		groupApi.GET("/messages", controller.FindMessages)
		groupApi.POST("/messages", controller.CreateMessage)
		groupApi.GET("/messages/:id", controller.FindMessage)
		groupApi.PATCH("/messages/:id", controller.UpdateMessage)
		groupApi.DELETE("/messages/:id", controller.DeleteMessage)
	}
}
