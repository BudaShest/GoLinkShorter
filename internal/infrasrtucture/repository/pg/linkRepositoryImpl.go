package pg

import (
	"context"
	"github.com/jackc/pgx/v4"
	"shorter/internal/entity"
)

type LinkRepository struct {
	db *pgx.Conn
}

var linkRepositoryInstance *LinkRepository

func GetLinkRepositoryInstance(db *pgx.Conn) *LinkRepository {
	if linkRepositoryInstance == nil {
		linkRepositoryInstance = &LinkRepository{db}
	}
	return linkRepositoryInstance
}

// todo embed shit
var createLinkQuery string

func (r *LinkRepository) CreateLink(ctx context.Context, link *entity.Link) error {
	model, err := toPgLink(link)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, createLinkQuery, model.Label, model.FullLink, model.ShortLink)

	if err != nil {
		return err
	}

	return nil
}

var getLinkUserQuery string

func (r *LinkRepository) GetLink(ctx context.Context, linkId int) (*entity.Link, error) {
	row := r.db.QueryRow(ctx, getLinkUserQuery, linkId)

	model := &Link{}
	err := row.Scan(&model.Id, &model.Label, &model.FullLink, &model.ShortLink, &model.CreatedAt, &model.UpdatedAt)
	if err != nil {
		return nil, err
	}

	var entityLink *entity.Link
	entityLink, err = toEntityLink(model)
	return entityLink, nil
}

var fetchLinkUserQuery string

func (r *LinkRepository) FetchLink(ctx context.Context, shortLink string) (*entity.Link, error) {
	row := r.db.QueryRow(ctx, fetchLinkUserQuery, shortLink)

	model := &Link{}
	err := row.Scan(&model.Id, &model.Label, &model.FullLink, &model.ShortLink, &model.CreatedAt, &model.UpdatedAt)
	if err != nil {
		return nil, err
	}

	var entityLink *entity.Link
	entityLink, err = toEntityLink(model)
	return entityLink, nil
}