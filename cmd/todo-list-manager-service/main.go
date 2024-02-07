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
	"os"
	"os/signal"
	"syscall"
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

	accountsRepo := repository.PostgresAccountRepository{DB: db}
	accountsService := service.PostgresAccountsService{Repo: &accountsRepo}
	accountController := controllers.AccountsController{Service: &accountsService}

	listRepo := repository.PostgresTodoListRepository{DB: db}
	listService := service.PostgresTodoListService{Repo: &listRepo}
	listController := controllers.TodoListController{Service: &listService}

	jwtService := service.JwtService{}
	authController := &controllers.AuthController{JwtService: &jwtService, AccountsService: &accountsService}

	router := mux.NewRouter()
	authenticatedRoute := router.PathPrefix("/secure").Subrouter()
	authenticatedRoute.Use(authController.AuthenticateMiddleware)

	// Available routes
	router.HandleFunc("/ping", listController.Ping)
	router.HandleFunc("/register", accountController.Register)              // register new account
	router.HandleFunc("/authenticate", authController.Authenticate)         // authenticate and retrieve short-lived token
	authenticatedRoute.HandleFunc("/create", listController.CreateTodoList) // creates to-do-list

	// close database connection on SIGTERM
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-signalChannel
		fmt.Printf("Received signal %v. Shutting down...\n", sig)

		err := db.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v\n", err)
		}

		os.Exit(0)
	}()

	fmt.Println("Server listening on port 9000...")
	err = http.ListenAndServe(":9000", router)
	if err != nil {
		log.Fatal("Failed to handle request: ", err)
	}
}
