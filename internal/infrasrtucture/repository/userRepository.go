package repository

import (
	"context"
	"shorter/internal/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, username string) (*entity.User, error)
}
