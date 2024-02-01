package service

import (
	"github.com/google/uuid"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/models"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/repository"
)

type TodoListService interface {
	CreateTodoList(userId string, name string) (string, error)
}

type PostgresTodoListService struct {
	Repo repository.TodoListRepository
}

func (s *PostgresTodoListService) CreateTodoList(userId string, name string) (string, error) {
	listId := uuid.NewString()
	todoList := models.NewTodoList(listId, userId, name)

	err := s.Repo.CreateTodoList(&todoList)
	if err != nil {
		//FIXME log error here
		return listId, err
	}
	return listId, nil
}
