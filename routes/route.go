package routes

import (
	"aprian1337/thukul-service/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func V1() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	api := e.Group("api/v1")
	api.GET("/salaries", controllers.GetSalariesController)
	api.POST("/salaries", controllers.CreateSalariesController)
	//api.POST("/login", controllers.LoginUsersController)
	api.GET("/users", controllers.GetUsersController)
	api.POST("/users", controllers.CreateUsersController)
	api.POST("/auth/login", controllers.LoginUsersController)
	return e
}