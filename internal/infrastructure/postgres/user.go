package postgres

import (
	"context"
	"x_golang_api/internal/domain/model"
	"x_golang_api/internal/domain/repository"
	"x_golang_api/internal/infrastructure/postgres/gen"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
	q *gen.Queries
}

func NewUserRepository(db *pgxpool.Pool) repository.UserRepository {
	return &userRepository{
		db: db,
		q:  gen.New(db),
	}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
    params := gen.InsertUserParams  {
		Email: user.Email,
		PasswordHash: user.HashedPassword,
	}

    inserted, err := ur.q.InsertUser(ctx, params)
	if err != nil {
		return nil, err
	}
    return &model.User{
        UserID:         inserted.UserID,
        Email:          inserted.Email,
        HashedPassword: inserted.PasswordHash,
        CreatedAt:      inserted.CreatedAt.Time,
    }, nil
}

func (ur *userRepository) ActivateUser(ctx context.Context, userID int32) error {
	err := ur.q.ActivateUser(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := ur.q.FindByEmail(ctx, email)
	if err != nil {
		return &model.User{}, err
	}
	return &model.User{
        UserID:         user.UserID,
        Email:          user.Email,
        HashedPassword: user.PasswordHash,
        CreatedAt:      user.CreatedAt.Time,
    }, nil
}

