package postgres

import (
	"fmt"
	"github.com/DeMarDeXis/VProj/internal/model"
	"github.com/jmoiron/sqlx"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user model.User) (int, error) {
	var id int
	q := fmt.Sprintf("INSERT INTO %s (name, username) values ($1, $2) RETURNING id", usersTable)
	row := r.db.QueryRow(q, user.Name, user.Username)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Auth) GetUser(username string) (model.User, error) {
	var user model.User
	q := fmt.Sprintf("SELECT id FROM %s WHERE username=$1", usersTable)
	err := r.db.Get(&user, q, username)

	return user, err
}
