package pg

import (
	"errors"
	"shorter/internal/entity"
	"strconv"
)

//TODO теги bson (тока для pgx)?
type User struct {
	Id       int
	Login    string
	Password string
}

func toPgUser(u *entity.User) (*User, error) {
	userId, err := strconv.Atoi(u.Id)

	if err != nil {
		return nil, err
	}

	return &User{
		Id:       userId,
		Login:    u.Login,
		Password: u.Password,
	}, nil
}

//TODO потом бы на нормальные UID перейти
func toEntity(u *User) (*entity.User, error) {
	userId := strconv.Itoa(u.Id)

	if userId == "" {
		return nil, errors.New("id of entity can not be nil")
	}

	return &entity.User{
		Id:       userId,
		Login:    u.Login,
		Password: u.Password,
	}, nil
}
