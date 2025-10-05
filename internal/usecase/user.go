package usecase

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"x_golang_api/internal/domain/model"
	"x_golang_api/internal/domain/repository"
	"x_golang_api/internal/domain/service"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type UserService interface {
    SignUp(c context.Context, email, password string) (*model.User, string, error)
	Login(ctx context.Context, email, password string) (string, error)
}

type userService struct {
	userRepo         repository.UserRepository
	tokenRepo        repository.TokenRepository
	passwordHasher   service.PasswordHasher
}

func NewUserService(
	ur             repository.UserRepository, 
	passwordHasher service.PasswordHasher,
	tr             repository.TokenRepository,
) UserService {
	return &userService{
		userRepo:       ur,
		passwordHasher: passwordHasher, 
		tokenRepo:      tr,
	}
}

func (uu *userService) SignUp(ctx context.Context, email, password string) (*model.User, string, error) {
	hash, err := uu.passwordHasher.HashPassword(password)
	if err != nil {
		return nil, "", err
	}

    user := &model.User{
		Email:          email,
		HashedPassword: string(hash),
	}
	createdUser, err := uu.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, "", err
	}
	// メール認証token生成と、redisへの保存依頼
	uuid := uuid.New()
	token := fmt.Sprintf("activation:%s", uuid.String())

	userID := createdUser.UserID
	expiration := 1 * time.Hour

	log.Printf("INFO: Calling TokenRepository with UserID: %d", createdUser.UserID)
	err = uu.tokenRepo.SetToken(ctx, token, userID, expiration)
	if err != nil {
		log.Printf("ERROR: Usecase received an error from TokenRepository: %v", err)
		return nil, "", err
	}


	log.Println("INFO: Usecase completed SignUp successfully.") 
	return createdUser, uuid.String(), nil
}

func (uu *userService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := uu.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = uu.passwordHasher.CompareHashAndPassword(ctx, user.HashedPassword, password)
	if err != nil {
		return "", err
	}
	// ペイロードとヘッダー
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	// 署名
	jwtSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}