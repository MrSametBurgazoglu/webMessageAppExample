package router

import "github.com/gin-gonic/gin"
import "webMessageApp/controller"

func SetPersonRouter(engine *gin.Engine) {
	peopleApi := engine.Group("/person")
	{
		peopleApi.GET("/people", controller.FindPeople)
		peopleApi.POST("/people", controller.CreatePerson)
		peopleApi.GET("/people/:id", controller.FindPerson)
		peopleApi.GET("/people_chats/:id", controller.GetPersonChats)
		peopleApi.PATCH("/people/:id", controller.UpdatePerson)
		peopleApi.DELETE("/people/:id", controller.DeletePerson)
	}
}
