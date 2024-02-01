package models

type Account struct {
	Id           string
	UserId       string
	UserName     string
	PasswordHash string
}

func NewAccount(id string, userId string, username string, passwordHash string) Account {
	account := Account{id, userId, username, passwordHash}
	return account
}
