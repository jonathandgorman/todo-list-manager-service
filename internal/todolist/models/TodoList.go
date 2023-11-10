package models

type TodoList struct {
	Id     string
	UserId string
}

func NewTodoList(id string, userId string) TodoList {
	todoList := TodoList{id, userId}
	return todoList
}
