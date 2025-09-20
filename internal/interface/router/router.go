package router

import (
	"x_golang_api/internal/interface/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(uc handler.UserHandler) *gin.Engine {
	router := gin.Default()
	router.POST("/signup", uc.SignUp)
	router.PUT("/activate/:user_id", uc.SignUp)
	return router
}