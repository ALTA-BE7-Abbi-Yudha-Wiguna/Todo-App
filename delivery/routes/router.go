package routes

import (
	"github.com/labstack/echo/v4"
	"todoListApp/delivery/handler/auth"
	"todoListApp/delivery/handler/project"
	"todoListApp/delivery/handler/todo"
	"todoListApp/delivery/handler/user"
	"todoListApp/delivery/middlewares"
)

func RegisterAuthPath(e *echo.Echo, ah *auth.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}

func RegisterPath(e *echo.Echo, uh *user.UserHandler, th *todo.TodoHandler, ph *project.ProjectHandler) {

	e.GET("/users", uh.GetAllHandler(), middlewares.JWTMiddleware())
	e.GET("/users/profile", uh.GetUserById(), middlewares.JWTMiddleware())
	e.POST("/users", uh.CreateUser())
	e.DELETE("/users", uh.DeleteUser(), middlewares.JWTMiddleware())
	e.PUT("/users", uh.UpdateUser(), middlewares.JWTMiddleware())

	e.GET("/todos", th.GetAllTodoHandler(), middlewares.JWTMiddleware())
	e.GET("/todos/:id", th.GetTodoById(), middlewares.JWTMiddleware())
	e.POST("/todos", th.CreateTodo(), middlewares.JWTMiddleware())
	e.DELETE("/todos/:id", th.DeleteTodo(), middlewares.JWTMiddleware())
	e.PUT("/todos/:id", th.UpdateTodo(), middlewares.JWTMiddleware())
	e.GET("/todos/profile", th.GetAllTodosByIdUser(), middlewares.JWTMiddleware())

	e.POST("/project", ph.CreateProject(), middlewares.JWTMiddleware())
	e.DELETE("/project/:id", ph.DeleteProject(), middlewares.JWTMiddleware())
	e.PUT("/project/:id", ph.UpdateProject(), middlewares.JWTMiddleware())
	e.GET("/project/profile", ph.GetAllProjectByUserId(), middlewares.JWTMiddleware())

	e.POST("/complete", th.CompleteTodo(), middlewares.JWTMiddleware())
	e.POST("/reopen", th.ReOpen(), middlewares.JWTMiddleware())
}
