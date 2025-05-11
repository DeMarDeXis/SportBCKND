package model

import (
	"time"
)

type TodoList struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"description" db:"description"`
	DoeDate     time.Time `json:"doe_date" db:"doe_date"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

//type UsersList struct {
//	ID     int
//	UserID int
//	ListID int
//}
//
//type UpdateListInput struct {
//	Title       *string `json:"title"`
//	Description *string `json:"description"`
//	DoeDate     *string `json:"doe_date"`
//}
//
//func (l UpdateListInput) Validate() error {
//	if l.Title == nil && l.Description == nil && l.DoeDate == nil {
//		return errors.New("update struct has no values")
//	}
//
//	return nil
//}
