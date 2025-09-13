package handler

import (
	"net/http"
	"x_golang_api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)


type UserHandler interface {
	SignUp(c *gin.Context) 
}

type SignUpRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"`
}

type UserResponse struct {
	UserID       int32              `json:"user_id"`
	Email        string             `json:"email"`
	CreatedAt    pgtype.Timestamp   `json:"created_at"`
}

type userHandler struct {
	us service.UserService
}

func NewUserHandler(uu service.UserService) UserHandler {
	return &userHandler{uu}
}

func (uh *userHandler) SignUp(c *gin.Context) {
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser, err := uh.us.SignUp(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := UserResponse{
		UserID: createdUser.UserID,
		Email:  createdUser.Email,
		CreatedAt: createdUser.CreatedAt,
	}
	c.JSON(http.StatusCreated, res)
}
