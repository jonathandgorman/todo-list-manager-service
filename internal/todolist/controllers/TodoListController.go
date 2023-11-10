package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/service"
	"net/http"
)

type Controller struct {
	Service service.TodoListService
}

func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("PONG")
}

func (c *Controller) CreateTodoList(w http.ResponseWriter, r *http.Request) {
	//FIXME add some sort of basic auth, for now use a dummy user ID
	userId := uuid.NewString()

	id, err := c.Service.CreateTodoList(userId)
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(w).Encode(id)
	if err != nil {
		panic(err)
	}
}
