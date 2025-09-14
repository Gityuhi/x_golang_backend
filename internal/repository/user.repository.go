package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"x_golang_api/internal/gen"
)


type UserRepository interface {
	CreateUser(ctx context.Context, email, hashedPassword string) (gen.User, error) 
}

type userRepository struct {
	db *pgxpool.Pool
	q *gen.Queries
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{
		db: db,
		q:  gen.New(db),
	}
}

func (ur *userRepository) CreateUser(ctx context.Context, email, hashedPassword string) (gen.User, error) {
	params := gen.InsertUserParams  {
		Email: email,
		PasswordHash: hashedPassword,
	}

	user, err := ur.q.InsertUser(ctx, params)
	if err != nil {
		return gen.User{}, err
	}
	return user, nil
}
