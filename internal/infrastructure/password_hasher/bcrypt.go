package bcrypt

import (
	"context"
	"x_golang_api/internal/domain/service"

	"golang.org/x/crypto/bcrypt"
)


type BcryptHasher struct {}

func NewBcryptHasher() service.PasswordHasher {
	return &BcryptHasher{}
}

func (h *BcryptHasher) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func (h *BcryptHasher) CompareHashAndPassword(ctx context.Context, HashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(password))
}