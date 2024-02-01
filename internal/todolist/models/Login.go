package models

type Login struct {
	Username string
	Password string
}

func NewLogin(username string, password string) Login {
	login := Login{username, password}
	return login
}
