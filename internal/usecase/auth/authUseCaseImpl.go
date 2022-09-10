package auth

import (
	"context"
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

}

func (u *UseCaseImpl) SignIn(ctx context.Context, username, password string) (string, error) {

}

func (u *UseCaseImpl) ParseToken(ctx context.Context, accessToken string) (*entity.User, error) {

}