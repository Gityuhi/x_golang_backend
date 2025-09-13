package router

import (
	"x_golang_api/internal/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(uh handler.UserHandler) *gin.Engine {
	router := gin.Default()
	router.POST("/signup", uh.SignUp)
	return router
}