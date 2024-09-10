package postgres

import (
	"fmt"
	"github.com/DeMarDeXis/VProj/internal/model"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"strings"
)

type TodoList struct {
	db     *sqlx.DB
	logger *slog.Logger
}

func NewTodoList(db *sqlx.DB, log *slog.Logger) *TodoList {
	return &TodoList{
		db:     db,
		logger: log,
	}
}

func (t *TodoList) Create(userID int, list model.TodoList) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		t.logger.Error("failed to begin tx", slog.String("error", err.Error()))
		return 0, err
	}

	var id int
	q := fmt.Sprintf("INSERT INTO %s (title, description, doe_date, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id", todoListsTable)
	row := tx.QueryRow(q, list.Title, list.Description, list.DoeDate)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		t.logger.Error("failed to scan id", slog.String("error", err.Error()))
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userID, id)
	if err != nil {
		tx.Rollback()
		t.logger.Error("failed to exec create users list query", slog.String("error", err.Error()))
		return 0, err
	}

	return id, tx.Commit()
}

func (t *TodoList) GetAll(userID int) ([]model.TodoList, error) {
	var lists []model.TodoList
	q := fmt.Sprintf(
		`SELECT tl.id, tl.title, tl.description, tl.doe_date, tl.created_at, tl.updated_at 
				FROM %s tl 
				INNER JOIN %s ul 
				ON tl.id = ul.list_id 
				WHERE ul.user_id = $1`,
		todoListsTable, usersListsTable)
	err := t.db.Select(&lists, q, userID)
	if err != nil {
		t.logger.Error("failed to select lists", slog.String("error", err.Error()))
		return nil, err
	}

	return lists, nil
}

func (t *TodoList) GetByID(userID, listID int) (model.TodoList, error) {
	var list model.TodoList

	q := fmt.Sprintf(`
		SELECT tl.id, tl.title, tl.description
		FROM %s tl 
		INNER JOIN %s ul 
		ON tl.id = ul.list_id 
		WHERE ul.user_id = $1 
		AND ul.list_id = $2`,
		todoListsTable, usersListsTable)

	err := t.db.Get(&list, q, userID, listID)
	if err != nil {
		t.logger.Error("failed to get list by id", slog.String("error", err.Error()))
		return model.TodoList{}, err
	}

	return list, nil
}

func (t *TodoList) Delete(userID, listID int) error {
	q := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable, usersListsTable)
	_, err := t.db.Exec(q, userID, listID)
	if err != nil {
		t.logger.Error("failed to delete list", slog.String("error", err.Error()))
		return err
	}

	return nil
}

func (t *TodoList) Update(userID, listID int, input model.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, *input.Description)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	q := fmt.Sprintf(`UPDATE %s tl 
							SET %s 
							FROM %s ul 
							WHERE tl.id = ul.list_id 
							AND ul.list_id = $%d AND ul.user_id = $%d`,
		todoListsTable, setQuery, usersListsTable, argID, argID+1)

	t.logger.Debug("Update query", slog.String("query", q))
	t.logger.Debug("Update args", slog.Any("args", args))

	_, err := t.db.Exec(q, args...)
	if err != nil {
		t.logger.Error("failed to update list", slog.String("error", err.Error()))
		return err
	}

	return err
}
