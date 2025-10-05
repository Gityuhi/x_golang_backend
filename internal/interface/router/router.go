package router

import (
	"x_golang_api/internal/interface/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler handler.UserHandler, activateHandler handler.ActivateUser) *gin.Engine {
	router := gin.Default()
	router.POST("/signup", userHandler.SignUp)
	router.POST("/login", userHandler.Login)
	router.GET("/activate", func(c *gin.Context) {
		err := activateHandler.Activate(c)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "アカウントが正常に有効化されました"})
	})
	return router
}