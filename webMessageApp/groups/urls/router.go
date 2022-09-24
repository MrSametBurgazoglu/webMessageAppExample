package urls

import (
	"github.com/gin-gonic/gin"
	"webMessageApp/groups/controller"
)

func SetRouter(engine *gin.Engine) {
	groupApi := engine.Group("/group")
	{
		groupApi.GET("/groups", controller.FindGroups)
		groupApi.POST("/groups", controller.CreateGroup)
		groupApi.GET("/groups/:id", controller.FindGroup)
		groupApi.PATCH("/groups/:id", controller.UpdateGroup)
		groupApi.DELETE("/groups/:id", controller.DeleteGroup)
	}
}
