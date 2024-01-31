package models

type TodoList struct {
	Id     string
	UserId string
	Name   string
}

func NewTodoList(id string, userId string, name string) TodoList {
	todoList := TodoList{id, userId, name}
	return todoList
}
