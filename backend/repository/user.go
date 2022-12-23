package repository

import "backend/model"

type User interface {
	GetUser(id int) (*model.User, error)
	PostUser(user *model.User) (int, error)
}
