package redis

import (
	"context"
	"fmt"
	"log"
	"time"
	"x_golang_api/internal/domain/repository"

	"github.com/redis/go-redis/v9"
)


type tokenRepository struct {
	client *redis.Client
}

func NewTokenRepository(client *redis.Client) repository.TokenRepository {
	return &tokenRepository{client: client}
}

func (t *tokenRepository) SetToken(ctx context.Context, token string, userID int32, expiration time.Duration) error {
	log.Printf("INFO: Setting token to Redis. Key: %s, UserID: %d", token, userID) 
	err := t.client.Set(ctx, token, userID, expiration).Err()
	if err != nil {
		log.Printf("ERROR: Failed to set token to Redis: %v", err)
		return err
	}
	log.Println("INFO: Successfully set token to Redis.")
	return nil
}

func (t *tokenRepository) ConsumeToken(ctx context.Context, token string) (UserID int32, err error) {
	activateToken := fmt.Sprintf("activation:%s", token)
	// keyからvalueを取得して、keyを削除する
	cmd := t.client.GetDel(ctx, activateToken)
	if cmd.Err() != nil {
		fmt.Println("tokenが見つかりませんでした。")
		return 0, cmd.Err()
	}

	userID64, err := cmd.Int64()
	if err != nil {
		return 0, err
	}
	userID := int32(userID64)

	return userID, nil
}
