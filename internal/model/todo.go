package model

import (
	"encoding/json"
	"errors"
	"time"
)

type TodoList struct {
	ID          int             `json:"id" db:"id"`
	Title       string          `json:"title" db:"title" binding:"required"`
	Description string          `json:"description" db:"description"`
	DoeDate     json.RawMessage `json:"doe_date" db:"doe_date"`
	CreatedAt   time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" db:"updated_at"`
}

type DateInput struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type UsersList struct {
	ID     int
	UserID int
	ListID int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	DoeDate     *string `json:"doe_date"`
}

func (l UpdateListInput) Validate() error {
	if l.Title == nil && l.Description == nil && l.DoeDate == nil {
		return errors.New("update struct has no values")
	}

	return nil
}
