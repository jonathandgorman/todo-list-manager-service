package controllers

import (
	"encoding/json"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/service"
	"io"
	"net/http"
)

type AccountsController struct {
	Service *service.PostgresAccountsService
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *AccountsController) Register(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		println("Failed to read register request")
		return
	}

	var registerRequest RegisterRequest
	if err := json.Unmarshal(bodyBytes, &registerRequest); err != nil {
		http.Error(w, "Error decoding JSON body", http.StatusBadRequest)
		return
	}

	exists, err := c.Service.UsernameExists(registerRequest.Username)
	if err != nil {
		http.Error(w, "Failed to create user account", http.StatusInternalServerError)
		return
	}

	if exists {
		http.Error(w, "That username is not available", http.StatusBadRequest)
		return
	}

	err = c.Service.Register(registerRequest.Username, registerRequest.Password)
	if err != nil {
		http.Error(w, "Failed to create user account", http.StatusInternalServerError)
		return
	}
}
