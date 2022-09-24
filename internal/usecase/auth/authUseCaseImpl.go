package auth

import (
	"context"
	"crypto/sha1"
	"fmt"
	"shorter/internal/entity"
	"shorter/internal/infrasrtucture/repository"
	"time"
)

type UseCaseImpl struct {
	userRepo        repository.UserRepository
	signingKey      []byte
	expiredDuration time.Duration
}

func NewAuthUseCase(userRepo repository.UserRepository, signingKey []byte, tokenTimeLife time.Duration) *UseCaseImpl {
	return &UseCaseImpl{
		userRepo:        userRepo,
		signingKey:      signingKey,
		expiredDuration: tokenTimeLife,
	}
}

func (u *UseCaseImpl) SingUp(ctx context.Context, username, password string) error {
	pwd := sha1.New()
	pwd.Write([]byte(password))

	//TODO ввалидация

	user := &entity.User{
		Login:    username,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}

	return u.userRepo.CreateUser(ctx, user)
}

func (u *UseCaseImpl) SignIn(ctx context.Context, username, password string) (string, error) {
	_, err := u.userRepo.GetUser(ctx, username)

	if err != nil {
		return "", err
	}
	return "", err
}

func (u *UseCaseImpl) ParseToken(ctx context.Context, accessToken string) (*entity.User, error) {
	return &entity.User{}, nil
}
