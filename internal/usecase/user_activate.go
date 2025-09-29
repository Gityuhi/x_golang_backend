package usecase

import (
	"context"
	"x_golang_api/internal/domain/repository"
)

type ActivateUser interface {
	UserActivate(ctx context.Context, token string) error
}


type userActivateService struct{
	// infra層のredis, postgres
	redis repository.TokenRepository
	postgres repository.UserRepository
}

func NewUserActivateService(
	redis repository.TokenRepository, 
	postgres repository.UserRepository,
) ActivateUser {
	return &userActivateService{
		redis:       redis,
		postgres:      postgres,
	}
}

func (a *userActivateService) UserActivate(ctx context.Context, token string) error {
	userID, err := a.redis.ConsumeToken(ctx, token)
	if err != nil {
		return err
	}

	err = a.postgres.ActivateUser(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}