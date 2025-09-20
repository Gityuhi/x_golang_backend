package handler

import (
	"net/http"
	"time"
	"x_golang_api/internal/usecase"

	"github.com/gin-gonic/gin"
)


type UserHandler interface {
	SignUp(c *gin.Context) 
}

type SignUpRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"`
}

type UserResponse struct {
    UserID       int32  `json:"user_id"`
    Email        string `json:"email"`
    CreatedAt    string `json:"created_at"`
}

type userHandler struct {
	uu usecase.UserService
}

func NewUserHandler(uu usecase.UserService) UserHandler {
	return &userHandler{uu}
}

func (uc *userHandler) SignUp(c *gin.Context) {
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser, err := uc.uu.SignUp(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := UserResponse{
		UserID: createdUser.UserID,
		Email:  createdUser.Email,
        CreatedAt: createdUser.CreatedAt.Format(time.RFC3339),
	}
	c.JSON(http.StatusCreated, res)
}
