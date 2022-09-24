package urls

import (
	"github.com/gin-gonic/gin"
	"webMessageApp/groups/controller"
)

func SetRouter(engine *gin.Engine) {
	peopleApi := engine.Group("/person")
	{
		peopleApi.GET("/people", controller.FindGroups)
		peopleApi.POST("/people", controller.CreateGroup)
		peopleApi.GET("/people/:id", controller.FindGroup)
		peopleApi.PATCH("/people/:id", controller.UpdateGroup)
		peopleApi.DELETE("/people/:id", controller.DeleteGroup)
	}
}
