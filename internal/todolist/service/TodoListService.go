package service

import (
	"github.com/google/uuid"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/models"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/repository"
)

type TodoListService interface {
	CreateTodoList(userId string) (string, error)
}

type Service struct {
	Repo repository.PostgresListRepository
}

func (s *Service) CreateTodoList(userId string) (string, error) {
	id := uuid.NewString()
	todoList := models.NewTodoList(id, userId)

	err := s.Repo.CreateTodoList(&todoList)
	if err != nil {
		//FIXME log error here
		return id, err
	}
	return id, nil
}
