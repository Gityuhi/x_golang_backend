package service

import "context"

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CompareHashAndPassword(ctx context.Context, HashedPassword, password string) error
}
