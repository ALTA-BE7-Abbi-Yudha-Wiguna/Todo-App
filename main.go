package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"todoListApp/config"
	"todoListApp/delivery/middlewares"
	"todoListApp/delivery/routes"
	"todoListApp/utils"

	authhandler "todoListApp/delivery/handler/auth"
	authrepo "todoListApp/repository/auth"
	authusecase "todoListApp/usecase/auth"

	userhandler "todoListApp/delivery/handler/user"
	userrepo "todoListApp/repository/user"
	userusecase "todoListApp/usecase/user"

	todohandler "todoListApp/delivery/handler/todo"
	todorepo "todoListApp/repository/todo"
	todousecase "todoListApp/usecase/todo"
)

func main() {
	configs := config.GetConfig()
	db := utils.InitDB(configs)

	userRepo := userrepo.NewUserRepository(db)
	userUseCase := userusecase.NewUserUseCase(userRepo)
	userHandler := userhandler.NewUserHandler(userUseCase)

	authRepo := authrepo.NewAuthRepository(db)
	authUseCase := authusecase.NewAuthUseCase(authRepo)
	authHandler := authhandler.NewAuthHandler(authUseCase)

	todoRepo := todorepo.NewTodoRepository(db)
	todoUse := todousecase.NewTodoUseCase(todoRepo)
	todoHandler := todohandler.NewTodoHandler(todoUse)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middlewares.CustomLogger())

	routes.RegisterAuthPath(e, authHandler)
	routes.RegisterPath(e, userHandler, todoHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%v", configs.Port)))

}
