package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/controllers"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/repository"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/service"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("application.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to read config: ", err)
	}

	db, err := repository.GetPostgresDatabase()
	if err != nil {
		log.Fatal("Error getting Postgres database: ", err)
	}

	listRepo := repository.Repository{DB: db}
	listService := service.Service{Repo: &listRepo}
	listController := controllers.Controller{Service: &listService}

	jwtService := &service.JwtService{}
	authController := &controllers.AuthController{JwtService: jwtService}

	router := mux.NewRouter()
	authenticatedRoute := router.PathPrefix("/secure").Subrouter()
	authenticatedRoute.Use(authController.AuthenticateMiddleware)

	// Available routes
	router.HandleFunc("/ping", listController.Ping)
	router.HandleFunc("/authenticate", authController.Authenticate)
	authenticatedRoute.HandleFunc("/create", listController.CreateTodoList)

	fmt.Println("Server listening on port 9000...")
	err = http.ListenAndServe(":9000", router)
	if err != nil {
		log.Fatal("Failed to handle request: ", err)
	}
}
