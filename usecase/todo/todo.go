package todo

import (
	"errors"
	"todoListApp/entities"
	"todoListApp/repository/todo"
)

type TodoUseCase struct {
	TodoRepository todo.TodoRepositoryInterface
}

func NewTodoUseCase(todoRepo todo.TodoRepositoryInterface) TodoUseCaseInterface {
	return &TodoUseCase{
		TodoRepository: todoRepo,
	}

}

func (tuc *TodoUseCase) GetAllTodos() ([]entities.Todo, error) {
	todos, err := tuc.TodoRepository.GetAllTodos()
	return todos, err
}

func (tuc *TodoUseCase) GetTodoById(id int) (entities.Todo, error) {
	todo, err := tuc.TodoRepository.GetTodoById(id)
	return todo, err
}

func (tuc *TodoUseCase) CreateTodo(todo entities.Todo) error {
	err := tuc.TodoRepository.CreateTodo(todo)
	return err
}

func (tuc *TodoUseCase) DeleteTodo(id, userID int) error {
	err := tuc.TodoRepository.DeleteTodo(id, userID)
	return err
}

func (tuc *TodoUseCase) UpdateTodo(id int, todo entities.Todo) error {
	err := tuc.TodoRepository.UpdateTodo(id, todo)
	return err
}

func (tuc *TodoUseCase) CompleteTodo(todo entities.Todo) error {

	//var user entities.User
	completetodo, err := tuc.TodoRepository.GetTodoById(int(todo.ID))

	if completetodo.Status == "Completed" || completetodo.UserID != todo.UserID {
		return errors.New("the task is complete")
	}

	completetodo.Status = "Completed"
	tuc.TodoRepository.CompleteTodo(completetodo)

	return err
}

func (tuc *TodoUseCase) ReOpen(todo entities.Todo) error {

	//var user entities.User
	reopen, err := tuc.TodoRepository.GetTodoById(int(todo.ID))

	if reopen.Status == "not completed" || reopen.UserID != todo.UserID {
		return errors.New("the task is complete")
	}

	reopen.Status = "not completed"
	tuc.TodoRepository.ReOpen(reopen)

	return err
}

func (tuc *TodoUseCase) GetAllTodosByIdUser(UserID int) ([]entities.Todo, error) {
	todos, err := tuc.TodoRepository.GetAllTodosByIdUser(UserID)
	return todos, err
}
