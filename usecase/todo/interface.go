package todo

import "todoListApp/entities"

type TodoUseCaseInterface interface {
	GetAllTodos() ([]entities.Todo, error)
	GetTodoById(id int) (entities.Todo, error)
	CreateTodo(todo entities.Todo) error
	DeleteTodo(id, userID int) error
	UpdateTodo(id int, Todo entities.Todo) error
	CompleteTodo(todo entities.Todo) error
	ReOpen(todo entities.Todo) error
	GetAllTodosByIdUser(UserID int) ([]entities.Todo, error)
}
