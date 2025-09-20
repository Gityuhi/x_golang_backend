package usecase

import (
	"context"
	"x_golang_api/internal/domain/model"
	"x_golang_api/internal/domain/repository"
	"x_golang_api/internal/domain/service"
)

type UserService interface {
    SignUp(c context.Context, email, password string) (*model.User, error)
}

type userService struct {
	ur               repository.UserRepository
	passwordHasher   service.PasswordHasher
}

func NewUserService(
	ur repository.UserRepository, 
	passwordHasher service.PasswordHasher,
) UserService {
	return &userService{
		ur:             ur,
		passwordHasher: passwordHasher, 
	}
}

func (uu *userService) SignUp(ctx context.Context, email, password string) (*model.User, error) {
	hash, err := uu.passwordHasher.HashPassword(password)
	if err != nil {
		return nil, err
	}

    user := &model.User{
		Email:          email,
		HashedPassword: string(hash),
	}
	createdUser, err := uu.ur.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}