package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"x_golang_api/internal/domain/service"
	"x_golang_api/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	uu         usecase.UserService
	mailsender service.EmailSender
}

func NewUserHandler(uu usecase.UserService, mailsender service.EmailSender) UserHandler {
	return &userHandler{
		uu: uu,
		mailsender: mailsender,
	}
}

func (uc *userHandler) SignUp(c *gin.Context) {
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser, uuid, err := uc.uu.SignUp(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// activateURLを作成し、emailとともにメソッドの引数に渡す。
	err = godotenv.Load()
	if err != nil {
		fmt.Println("envファイルの読み込みに失敗しました。")
	}
	url := fmt.Sprintf("%s/activate?token=%s", 
		os.Getenv("URL"),
		uuid)

	err = uc.mailsender.SendMail(c.Request.Context(), url, createdUser.Email)
	if err != nil {
		log.Printf("ERROR: failed to send email: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send confirmation email"})
		return
	}

	res := UserResponse{
		UserID: createdUser.UserID,
		Email:  createdUser.Email,
        CreatedAt: createdUser.CreatedAt.Format(time.RFC3339),
	}
	fmt.Println("メール送信されました。")
	c.JSON(http.StatusCreated, res)
}
