package service

import (
	"context"
	"x_golang_api/internal/gen"
	"x_golang_api/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(c context.Context, email, password string) (gen.User, error)
}

type userService struct {
	ur repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{ur}
}

func (us *userService) SignUp(ctx context.Context, email, password string) (gen.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return gen.User{}, err
	}
	createdUser, err := us.ur.CreateUser(ctx, email, string(hash))
	if err != nil {
		return gen.User{}, err
	}
	return createdUser, nil
}