package repository

import (
	"context"
	"time"
	"x_golang_api/internal/domain/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error) 
	ActivateUser(ctx context.Context, userID int32) error
}


type TokenRepository interface {
	SetToken(ctx context.Context, token string, userID int32, expiration time.Duration)  error
	ConsumeToken(ctx context.Context, token string) (UserID int32, err error)
}


