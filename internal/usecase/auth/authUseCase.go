package auth

import (
	"context"
	"shorter/internal/entity"
)

type UseCase interface {
	SignUp(ctx context.Context, username, password string) error
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*entity.User, error) //todo должно ли это быть тту вообще
}
