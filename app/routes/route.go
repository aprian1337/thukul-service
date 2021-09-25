package routes

import (
	"aprian1337/thukul-service/deliveries/salaries"
	"aprian1337/thukul-service/deliveries/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware    middleware.JWTConfig
	UserController   users.Controller
	SalaryController salaries.Controller
}

func (cl *ControllerList) RouteUsers(e *echo.Echo) {
	v1 := e.Group("api/v1/")
	//AUTH
	v1.POST("auth/login", cl.UserController.LoginUserController)

	//USERS
	v1.GET("users", cl.UserController.GetUsersController)
	v1.GET("users/:id", cl.UserController.GetDetailUserController)
	v1.POST("users", cl.UserController.CreateUserController, middleware.JWTWithConfig(cl.JWTMiddleware))
	v1.DELETE("users/:id", cl.UserController.DeleteUserController)
	v1.PUT("users/:id", cl.UserController.UpdateUserController)

	//SALARIES
	v1.GET("salaries", cl.SalaryController.GetSalariesController)
	v1.GET("salaries/:id", cl.SalaryController.GetSalaryByIdController)
	v1.POST("salaries", cl.SalaryController.CreateSalaryController)
	v1.PUT("salaries/:id", cl.SalaryController.UpdateSalaryController)
	v1.DELETE("salaries/:id", cl.SalaryController.DestroySalaryController)
}
