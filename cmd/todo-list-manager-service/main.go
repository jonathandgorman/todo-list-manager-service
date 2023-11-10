package main

import (
	"fmt"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/controllers"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/repository"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/service"
	"github.com/julienschmidt/httprouter"
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
	listService := service.Service{&listRepo}
	listController := controllers.Controller{&listService}

	router := httprouter.New()
	router.HandlerFunc("POST", "/createList/", listController.CreateTodoList)
	err = http.ListenAndServe(":9000", router)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening on port: 9000")
}
