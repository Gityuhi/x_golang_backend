package model

import (
	"context"
	"time"
)

type User struct {
    UserID         int32
    Email          string
    HashedPassword string
    CreatedAt      time.Time
}

type FindByEmail interface {
    FindByEmail(ctx context.Context, email string) (*User, error)
}