package routes

import (
	"aprian1337/thukul-service/app/middlewares"
	"aprian1337/thukul-service/deliveries/activities"
	"aprian1337/thukul-service/deliveries/coins"
	"aprian1337/thukul-service/deliveries/cryptos"
	"aprian1337/thukul-service/deliveries/favorites"
	"aprian1337/thukul-service/deliveries/payments"
	"aprian1337/thukul-service/deliveries/pockets"
	"aprian1337/thukul-service/deliveries/salaries"
	"aprian1337/thukul-service/deliveries/users"
	"aprian1337/thukul-service/deliveries/wishlists"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware   middlewares.MongoConfig
	JWTMiddleware      middleware.JWTConfig
	JWTMiddlewareAdmin middleware.JWTConfig
	UserController     users.Controller
	SalaryController   salaries.Controller
	PocketController   pockets.Controller
	ActivityController activities.Controller
	CoinController     coins.Controller
	WishlistController wishlists.Controller
	FavoriteController favorites.Controller
	CryptoController   cryptos.Controller
	PaymentController  payments.Controller
}

func (cl *ControllerList) RouteUsers(e *echo.Echo) {
	v1 := e.Group("api/v1/")
	cl.LoggerMiddleware.Start(e)
	//AUTH
	v1.POST("auth/login", cl.UserController.LoginUserController)

	//USERS
	//middleware.JWTWithConfig(cl.JWTMiddleware)
	v1.GET("users", cl.UserController.GetUsersController, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsAdmin)
	v1.GET("users/:userId", cl.UserController.GetDetailUserController, middleware.JWTWithConfig(cl.JWTMiddleware))
	v1.POST("users", cl.UserController.CreateUserController)
	v1.DELETE("users/:userId", cl.UserController.DeleteUserController, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsAdmin)
	v1.PUT("users/:userId", cl.UserController.UpdateUserController, middleware.JWTWithConfig(cl.JWTMiddleware))

	//SALARIES
	v1.GET("salaries", cl.SalaryController.GetSalariesController)
	v1.GET("salaries/:id", cl.SalaryController.GetSalaryByIdController)
	v1.POST("salaries", cl.SalaryController.CreateSalaryController)
	v1.PUT("salaries/:id", cl.SalaryController.UpdateSalaryController, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsAdmin)
	v1.DELETE("salaries/:id", cl.SalaryController.DestroySalaryController, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsAdmin)

	//POCKETS
	v1.GET("users/:userId/pockets", cl.PocketController.Get, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.GET("users/:userId/pockets/:pocketId", cl.PocketController.GetById, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.GET("users/:userId/pockets/:pocketId/total", cl.PocketController.Total, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.POST("users/:userId/pockets", cl.PocketController.Create, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.PUT("users/:userId/pockets/:pocketId", cl.PocketController.Update, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.DELETE("users/:userId/pockets/:pocketId", cl.PocketController.Destroy, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)

	//ACTIVITIES
	v1.GET("users/:userId/pockets/:idPocket/activities", cl.ActivityController.Get, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.GET("users/:userId/pockets/:idPocket/activities/:id", cl.ActivityController.GetById, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.POST("users/:userId/pockets/:idPocket/activities", cl.ActivityController.Create, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.PUT("users/:userId/pockets/:idPocket/activities/:id", cl.ActivityController.Update, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.DELETE("users/:userId/pockets/:idPocket/activities/:id", cl.ActivityController.Destroy, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)

	//WISHLISTS
	v1.GET("users/:userId/wishlists", cl.WishlistController.Get, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.GET("users/:userId/wishlists/:wishlistId", cl.WishlistController.GetById, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.POST("users/:userId/wishlists", cl.WishlistController.Create, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.PUT("users/:userId/wishlists/:wishlistId", cl.WishlistController.Update, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.DELETE("users/:userId/wishlists/:wishlistId", cl.WishlistController.Destroy, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)

	//FAVORITES
	v1.GET("users/:userId/favorites", cl.FavoriteController.Get, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.GET("users/:userId/favorites/:favId", cl.FavoriteController.GetById, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.POST("users/:userId/favorites", cl.FavoriteController.Create, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.DELETE("users/:userId/favorites/:favId", cl.FavoriteController.Destroy, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)

	//PAYMENTS
	v1.POST("payments/topup", cl.PaymentController.TopUp, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsAdmin)
	v1.POST("users/:userId/payments/buy", cl.PaymentController.Buy, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.POST("users/:userId/payments/sell", cl.PaymentController.Sell, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.GET("payments/confirm/:uuid/:encrypt", cl.PaymentController.Confirm)

	//CRYPTOS
	v1.GET("users/:userId/cryptos", cl.CryptoController.GetByUser, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)
	v1.GET("users/:userId/cryptos/:cryptoId", cl.CryptoController.GetDetail, middleware.JWTWithConfig(cl.JWTMiddleware), middlewares.IsUserId)

	//COINS
	v1.GET("coins", cl.CoinController.GetBySymbol)
	v1.GET("coins/all", cl.CoinController.GetAllSymbol)
}
