package port

import (
	"context"

	"github.com/OzkrOssa/radius-micro-users/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserById(ctx context.Context, id uint64) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	ListUsers(ctx context.Context, skip, limit uint64) ([]domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id uint64) error
}
