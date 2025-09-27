package handler

import (
	"x_golang_api/internal/usecase"

	"github.com/gin-gonic/gin"
)


type ActivateUser interface{
	Activate(c *gin.Context) error
}

type activateUser struct{
	ua usecase.ActivateUser
}


func NewActivateUser(ua usecase.ActivateUser) ActivateUser {
	return &activateUser{
		ua: ua,
	}
}

func (a *activateUser) Activate(c *gin.Context) error {
	token := c.Query("token")

	err := a.ua.UserActivate(c, token)
	if err != nil {
		return err
	}
	return nil
}