package usecase

import (
	"context"
	"fmt"
	"log"
	"time"
	"x_golang_api/internal/domain/model"
	"x_golang_api/internal/domain/repository"
	"x_golang_api/internal/domain/service"

	"github.com/google/uuid"
)

type UserService interface {
    SignUp(c context.Context, email, password string) (*model.User, string, error)
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