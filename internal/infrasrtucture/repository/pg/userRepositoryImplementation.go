package pg

import (
	"context"
	"github.com/jackc/pgx/v4"
	"shorter/internal/entity"
)

type UserRepository struct {
	db *pgx.Conn
}

var userRepositoryInstance *UserRepository

func GetUserRepositoryInstance(db *pgx.Conn) *UserRepository {
	if userRepositoryInstance == nil {
		userRepositoryInstance = &UserRepository{db}
	}
	return userRepositoryInstance
}

//todo embed shit
var createUserQuery string

func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	model, err := toPgUser(user)

	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, createUserQuery, model.Id, model.Password, model.Login)

	if err != nil {
		return err
	}

	return nil
}

var searchUserQuery string

func (r *UserRepository) GetUser(ctx context.Context, username string) (*entity.User, error) {
	var user *User
	row := r.db.QueryRow(ctx, searchUserQuery, username)

	err := row.Scan(&user.Id, &user.Login, &user.Password)

	if err != nil {
		return nil, err
	}

	entity, err := toEntity(user)

	return entity, nil
}
