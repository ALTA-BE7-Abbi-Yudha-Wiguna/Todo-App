package todo

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"todoListApp/delivery/helper"
	"todoListApp/delivery/middlewares"
	"todoListApp/entities"
	"todoListApp/usecase/todo"
)

type TodoHandler struct {
	todoUseCase todo.TodoUseCaseInterface
}

func NewTodoHandler(todoUseCase todo.TodoUseCaseInterface) *TodoHandler {
	return &TodoHandler{
		todoUseCase: todoUseCase,
	}
}

func (th *TodoHandler) GetAllTodoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		todos, err := th.todoUseCase.GetAllTodos()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all todos", todos))
	}
}

func (th *TodoHandler) GetTodoById() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))

		todos, err := th.todoUseCase.GetTodoById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get todo by id", todos))
	}
}

func (th *TodoHandler) CreateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		var todo entities.Todo
		c.Bind(&todo)
		id := middlewares.ExtractToken(c)
		todo.UserID = uint(id)
		err := th.todoUseCase.CreateTodo(todo)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to create todo"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success Create task"))
	}
}

func (th *TodoHandler) DeleteTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))
		var userID = middlewares.ExtractToken(c)

		err := th.todoUseCase.DeleteTodo(id, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success delete task"))
	}
}

func (th *TodoHandler) UpdateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))
		var todos entities.Todo
		c.Bind(&todos)

		err := th.todoUseCase.UpdateTodo(id, todos)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update task", todos))
	}
}

func (th *TodoHandler) CompleteTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		var todo entities.Todo

		c.Bind(&todo)
		id := middlewares.ExtractToken(c)
		todo.UserID = uint(id)
		err := th.todoUseCase.CompleteTodo(todo)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to complete task"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success complete task", todo))
	}
}

func (th *TodoHandler) ReOpen() echo.HandlerFunc {
	return func(c echo.Context) error {
		var todo entities.Todo

		c.Bind(&todo)
		id := middlewares.ExtractToken(c)
		todo.UserID = uint(id)
		err := th.todoUseCase.ReOpen(todo)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to complete task"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success reopen task", todo))
	}
}

func (th *TodoHandler) GetAllTodosByIdUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		var id = middlewares.ExtractToken(c)

		users, err := th.todoUseCase.GetAllTodosByIdUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get task by user id", users))
	}
}
