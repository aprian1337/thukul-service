package routes

import (
	"aprian1337/thukul-service/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func V1() *echo.Echo {
	e := echo.New()
	viper.SetConfigFile(`config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	jwt := middleware.JWT([]byte(viper.GetString("signed_string")))
	e.Pre(middleware.RemoveTrailingSlash())
	api := e.Group("api/v1")
	api.GET("/salaries", controllers.GetSalariesController)
	api.POST("/salaries", controllers.CreateSalariesController)
	//api.POST("/login", controllers.LoginUsersController)
	api.GET("/users", controllers.GetUsersController, jwt)
	api.POST("/users", controllers.CreateUsersController)
	api.POST("/auth/login", controllers.LoginUsersController)
	return e
}
