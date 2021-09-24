package routes

import (
	"aprian1337/thukul-service/deliveries/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserController users.UserController
}

func (cl *ControllerList) RouteUsers(e *echo.Echo) {
	v1 := e.Group("api/v1/")
	v1.GET("users", cl.UserController.GetUsersController)
	v1.POST("users", cl.UserController.CreateUserController)
	v1.POST("auth/login", cl.UserController.LoginUserController)
}
