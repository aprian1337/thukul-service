package routes

import (
	"aprian1337/thukul-service/app/middlewares"
	"aprian1337/thukul-service/deliveries/activities"
	"aprian1337/thukul-service/deliveries/coins"
	"aprian1337/thukul-service/deliveries/pockets"
	"aprian1337/thukul-service/deliveries/salaries"
	"aprian1337/thukul-service/deliveries/users"
	"aprian1337/thukul-service/deliveries/wishlists"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	MiddlewareConfig   middlewares.MiddlewareConfig
	JWTMiddleware      middleware.JWTConfig
	UserController     users.Controller
	SalaryController   salaries.Controller
	PocketController   pockets.Controller
	ActivityController activities.Controller
	CoinController     coins.Controller
	WishlistController wishlists.Controller
}

func (cl *ControllerList) RouteUsers(e *echo.Echo) {
	v1 := e.Group("api/v1/")
	cl.MiddlewareConfig.Start(e)
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

	//POCKETS
	v1.GET("pockets", cl.PocketController.Get)
	v1.GET("pockets/:id", cl.PocketController.GetById)
	v1.GET("pockets/:id/total", cl.PocketController.Total)
	v1.POST("pockets", cl.PocketController.Create)
	v1.PUT("pockets/:id", cl.PocketController.Update)
	v1.DELETE("pockets/:id", cl.PocketController.Destroy)

	//POCKETS
	v1.GET("pockets/:idPocket/activities", cl.ActivityController.Get)
	v1.GET("pockets/:idPocket/activities/:id", cl.ActivityController.GetById)
	v1.POST("pockets/:idPocket/activities", cl.ActivityController.Create)
	v1.PUT("pockets/:idPocket/activities/:id", cl.ActivityController.Update)
	v1.DELETE("pockets/:idPocket/activities/:id", cl.ActivityController.Destroy)

	//WISHLISTS

	v1.GET("users/:userId/wishlists", cl.WishlistController.Get)
	v1.GET("users/:userId/wishlists/:wishlistId", cl.WishlistController.GetById)
	v1.POST("users/:userId/wishlists", cl.WishlistController.Create)
	v1.PUT("users/:userId/wishlists/:wishlistId", cl.WishlistController.Update)
	v1.DELETE("users/:userId/wishlists/:wishlistId", cl.WishlistController.Destroy)

	//COINS
	v1.GET("coins", cl.CoinController.GetBySymbol)
}
