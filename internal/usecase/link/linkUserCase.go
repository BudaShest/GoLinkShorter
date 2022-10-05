package link

import (
	"context"
	"shorter/internal/entity"
)

type UseCase interface {
	SaveLink(ctx context.Context, link *entity.Link) error
	HasLink(ctx context.Context, linkId int) bool
	GetLink(ctx context.Context, linkId int) (*entity.Link, error)
	FetchLink(ctx context.Context, shortLink string) (*entity.Link, error)
	GetAllUserLinks(ctx context.Context, userId int) ([]*entity.Link, error)
	DeleteLink(ctx context.Context, linkId int) bool
}
