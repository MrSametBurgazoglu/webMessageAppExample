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
		peopleApi.GET("/people_add_friend_list/:id", controller.GetPersonAddFriendList)
		peopleApi.POST("/people_add_friend_list/:id", controller.PersonAddFriendList)
		peopleApi.GET("/people_create_group_list/:id", controller.GetChatCreateFriendList)
		peopleApi.PATCH("/people/:id", controller.UpdatePerson)
		peopleApi.DELETE("/people/:id", controller.DeletePerson)
	}
}
