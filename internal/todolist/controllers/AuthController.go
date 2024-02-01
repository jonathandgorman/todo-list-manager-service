package controllers

import (
	"encoding/json"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/models"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/service"
	"io"
	"net/http"
)

type AuthController struct {
	JwtService      *service.JwtService
	AccountsService *service.PostgresAccountsService
}

type AuthenticateBody struct {
	UsernameKey string `json:"username"`
	PasswordKey string `json:"password"`
}

func (c *AuthController) Authenticate(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		println("failed to read request")
		return
	}

	var requestBody AuthenticateBody
	if err := json.Unmarshal(bodyBytes, &requestBody); err != nil {
		http.Error(w, "Error decoding JSON body", http.StatusBadRequest)
		return
	}

	login := models.NewLogin(requestBody.UsernameKey, requestBody.PasswordKey)
	if len(login.Username) == 0 || len(login.Password) == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	exists, err := c.AccountsService.Exists(login.Username, login.Password)
	if err != nil {
		http.Error(w, "Something unexpected went wrong", http.StatusInternalServerError)
		return
	}

	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	jwtToken, _ := c.JwtService.GetToken(login.Username)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("token: " + jwtToken))
	if err != nil {
		return
	}
}

func (c *AuthController) AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		verified, err := c.JwtService.VerifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}

		if !verified {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r) // auth success, proceed to next handler
	})
}
