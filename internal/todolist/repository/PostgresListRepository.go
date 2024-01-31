package repository

import (
	"database/sql"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/models"
)

type PostgresListRepository interface {
	CreateTodoList(todoList *models.TodoList) error
}

type Repository struct {
	DB *sql.DB
}

func (r *Repository) CreateTodoList(list *models.TodoList) error {
	_, err := r.DB.Exec("INSERT INTO lists(id, user_id, name) VALUES ($1, $2, $3)", list.Id, list.UserId, list.Name)
	if err != nil {
		return err
	}
	return nil
}
