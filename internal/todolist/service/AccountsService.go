package service

import (
	"github.com/google/uuid"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/models"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
)

const HashCost = 5

type AccountsService interface {
	Register(username string, password string) error
	Exists(username string, password string) (bool, error)
	UsernameExists(username string) (bool, error)
}

type PostgresAccountsService struct {
	Repo repository.PostgresAccountRepository
}

func (s *PostgresAccountsService) SaltAndHashPassword(password []byte) ([]byte, error) {
	outputPassword, err := bcrypt.GenerateFromPassword(password, HashCost)
	if err != nil {
		return nil, err
	}
	return outputPassword, nil
}

func (s *PostgresAccountsService) CompareHashAndPasswordWithSalt(hashedPassword []byte, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresAccountsService) Register(username string, password string) error {
	accountId := uuid.NewString()
	userId := uuid.NewString()
	passwordBytes, err := s.SaltAndHashPassword([]byte(password))
	if err != nil {
		log.Fatalf("Failed to generate salted and hashed password")
		return err
	}
	hashedPassword := string(passwordBytes)
	account := models.NewAccount(accountId, userId, username, hashedPassword)

	err = s.Repo.Register(&account)
	if err != nil {
		log.Fatalf("Failed to register account")
		return err
	}
	return nil
}

func (s *PostgresAccountsService) Exists(username string, password string) (exists bool, err error) {
	account, err := s.Repo.GetAccount(username)
	if err != nil {
		log.Fatalf("Failed to check existence of account with username %s", username)
		return
	}

	if account == (models.Account{}) {
		return false, nil // account does not exist
	}

	// compare plaintext password with hashedPassword to determine if they are related
	maybeError := s.CompareHashAndPasswordWithSalt([]byte(account.PasswordHash), []byte(password))
	if maybeError != nil {
		return false, maybeError
	}
	return true, nil
}

func (s *PostgresAccountsService) UsernameExists(username string) (exists bool, err error) {
	account, err := s.Repo.GetAccount(username)
	if err != nil {
		log.Fatalf("Failed to check existence of account with username %s", username)
		return
	}

	if account == (models.Account{}) {
		return false, nil // account does not exist
	}
	return true, nil
}
