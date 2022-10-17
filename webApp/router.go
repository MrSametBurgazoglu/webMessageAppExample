package webApp

import (
	"github.com/gin-gonic/gin"
	"webMessageApp/router"
)

func SetRouter(engine *gin.Engine) {
	router.SetPersonRouter(engine)
	router.SetMessageRouter(engine)
	engine.GET("/", HomePage)
	engine.GET("/chat/:id", ChatPage)
	engine.GET("/login", LoginPage)
	engine.GET("/signup", SignupPage)
	engine.GET("/create_person", CreatePersonPage)
}
