package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"todoListApp/config"
	authhandler "todoListApp/delivery/handler/auth"
	"todoListApp/delivery/middlewares"
	"todoListApp/delivery/routes"
	authrepo "todoListApp/repository/auth"
	authusecase "todoListApp/usecase/auth"
	"todoListApp/utils"

	userhandler "todoListApp/delivery/handler/user"
	userrepo "todoListApp/repository/user"
	userusecase "todoListApp/usecase/user"
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

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middlewares.CustomLogger())

	routes.RegisterAuthPath(e, authHandler)
	routes.RegisterPath(e, userHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%v", configs.Port)))

}
