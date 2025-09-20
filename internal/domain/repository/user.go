package repository

import (
	"context"
	"x_golang_api/internal/domain/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error) 
}