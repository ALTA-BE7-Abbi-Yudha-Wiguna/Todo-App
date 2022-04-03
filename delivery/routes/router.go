package routes

import (
	"github.com/labstack/echo/v4"
	"todoListApp/delivery/handler/user"
	"todoListApp/delivery/middlewares"
)

func RegisterPath(e *echo.Echo, uh *user.UserHandler) {

	e.GET("/users", uh.GetAllHandler(), middlewares.JWTMiddleware())
	e.GET("/users/profile", uh.GetUserById(), middlewares.JWTMiddleware())
	e.POST("/users", uh.CreateUser())
	e.DELETE("/users", uh.DeleteUser(), middlewares.JWTMiddleware())
	e.PUT("/users", uh.UpdateUser(), middlewares.JWTMiddleware())

}
