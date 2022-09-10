package repository

import "shorter/internal/entity"

type LinkRepository interface {
	CreateLink() error
	GetLink() (*entity.Link, error)
}
