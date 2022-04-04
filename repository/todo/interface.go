package todo

import "todoListApp/entities"

type TodoRepositoryInterface interface {
	GetAllTodos() ([]entities.Todo, error)
	GetTodoById(id int) (entities.Todo, error)
	CreateTodo(book entities.Todo) error
	DeleteTodo(id, userID int) error
	UpdateTodo(id int, book entities.Todo) error
	CompleteTodo(todo entities.Todo) error
	ReOpen(todo entities.Todo) error
	GetAllTodosByIdUser(UserID int) ([]entities.Todo, error)
}
