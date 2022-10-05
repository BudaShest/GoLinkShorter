package pg

import "shorter/internal/entity"

type Link struct {
	Id        string
	Label     string
	FullLink  string
	ShortLink string
	CreatedAt int //timestamp
	UpdatedAt int //timestamp
}

func toPgLink(l *entity.Link) (*Link, error) {
	return &Link{
		Id:        l.Id,
		Label:     l.Label,
		FullLink:  l.FullLink,
		ShortLink: l.ShortLink,
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
	}, nil
}

// TODO потом бы на нормальные UID перейти
func toEntityLink(l *Link) (*entity.Link, error) {
	return &entity.Link{
		Id:        l.Id,
		Label:     l.Label,
		FullLink:  l.FullLink,
		ShortLink: l.ShortLink,
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
	}, nil
}
