package app

import (
	"github.com/gin-gonic/gin"
	groupRouter "webMessageApp/groups/urls"
	messageRouter "webMessageApp/messages/urls"
	peopleRouter "webMessageApp/people/urls"
)

func SetRouter(engine *gin.Engine) {
	groupRouter.SetRouter(engine)
	messageRouter.SetRouter(engine)
	peopleRouter.SetRouter(engine)
	engine.GET("/", HomePage)
	engine.GET("/chat/:id", ChatPage)
}
