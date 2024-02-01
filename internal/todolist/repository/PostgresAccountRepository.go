package repository

import (
	"database/sql"
	"github.com/jonathandgorman/todo-list-manager-service/internal/todolist/models"
)

type AccountRepository interface {
	Register(todoList *models.Account) error
	GetAccount(username string, password string) (models.Account, bool, error)
}

type PostgresAccountRepository struct {
	DB *sql.DB
}

func (r *PostgresAccountRepository) Register(account *models.Account) error {
	_, err := r.DB.Exec("INSERT INTO accounts(account_id, user_id, username, password_hash) VALUES ($1, $2, $3, $4)", account.Id, account.UserId, account.UserName, account.PasswordHash)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresAccountRepository) GetAccount(username string) (models.Account, error) {
	var account models.Account
	row := r.DB.QueryRow("SELECT account_id, user_id, username, password_hash FROM accounts WHERE username = $1", username)

	err := row.Scan(&account.Id, &account.UserId, &account.UserName, &account.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Account{}, nil
		}
		return models.Account{}, err
	}
	return account, nil
}
