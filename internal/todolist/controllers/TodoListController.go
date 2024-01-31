package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/service"
	"io"
	"net/http"
)

type Controller struct {
	Service service.TodoListService
}

type CreateRequestBody struct {
	NameKey string `json:"name"`
}

func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode("PONG")
	if err != nil {
		return
	}
}

func (c *Controller) CreateTodoList(w http.ResponseWriter, r *http.Request) {

	userId := uuid.NewString() // random UUID
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var requestBody CreateRequestBody
	if err := json.Unmarshal(bodyBytes, &requestBody); err != nil {
		http.Error(w, "Error decoding JSON body", http.StatusBadRequest)
		return
	}

	listId, err := c.Service.CreateTodoList(userId, requestBody.NameKey)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(listId)
	if err != nil {
		panic(err)
	}
}
