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

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middlewares.CustomLogger())

	routes.RegisterPath(e, userHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%v", configs.Port)))

}
