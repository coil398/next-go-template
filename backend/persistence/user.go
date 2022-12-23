package persistence

import (
	"backend/model"
	"backend/repository"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type UserStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) repository.User {
	return &UserStore{db}
}

func (u *UserStore) GetUser(id int) (*model.User, error) {
	var user model.User
	row := u.db.QueryRowx("SELECT id, name FROM user WHERE id = ?", id)

	if err := row.StructScan(&user); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserStore) PostUser(user *model.User) (int, error) {
	res, err := u.db.Exec("INSERT INTO user (name) VALUES (?)", user.Name)

	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}
