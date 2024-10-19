package model

import "errors"

type User struct {
	ID       int    `json:"-" db:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u User) Validate() error {
	if u.Name == "" || u.Username == "" || u.Password == "" {
		return errors.New("invalid input data")
	}

	return nil
}
