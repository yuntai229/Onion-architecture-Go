package presentation

import (
	domain "onion-architecrure-go/domain/entity"
	handler "onion-architecrure-go/presentation/api/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(userApp domain.UserApp) *gin.Engine {
	router := gin.Default()

	homeHandler := handler.NewHomeController()
	userHandler := handler.NewUserController(userApp)

	router.GET("/ping", homeHandler.Ping)
	router.POST("/user/signup", userHandler.Signup)

	return router
}
