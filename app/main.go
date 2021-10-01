package main

import (
	"aprian1337/thukul-service/app/middlewares"
	"aprian1337/thukul-service/app/routes"
	_usersUsecase "aprian1337/thukul-service/business/users"
	_usersDelivery "aprian1337/thukul-service/deliveries/users"
	"aprian1337/thukul-service/helpers/constants"
	_transactionHistoryDb "aprian1337/thukul-service/repository/databases/transactions"
	_usersDb "aprian1337/thukul-service/repository/databases/users"

	_coinmarketRepo "aprian1337/thukul-service/repository/thirdparties/coinmarket"

	_salaryUsecase "aprian1337/thukul-service/business/salaries"
	_salaryDelivery "aprian1337/thukul-service/deliveries/salaries"
	_salaryDb "aprian1337/thukul-service/repository/databases/salaries"

	_pocketUsecase "aprian1337/thukul-service/business/pockets"
	_pocketDelivery "aprian1337/thukul-service/deliveries/pockets"
	_pocketDb "aprian1337/thukul-service/repository/databases/pockets"

	_activityUsecase "aprian1337/thukul-service/business/activities"
	_activityDelivery "aprian1337/thukul-service/deliveries/activities"
	_activityDb "aprian1337/thukul-service/repository/databases/activities"

	_coinUsecase "aprian1337/thukul-service/business/coins"
	_coinDelivery "aprian1337/thukul-service/deliveries/coins"
	_coinDb "aprian1337/thukul-service/repository/databases/coins"

	_wishlistUsecase "aprian1337/thukul-service/business/wishlists"
	_wishlistDelivery "aprian1337/thukul-service/deliveries/wishlists"
	_wishlistDb "aprian1337/thukul-service/repository/databases/wishlists"

	_favoriteUsecase "aprian1337/thukul-service/business/favorites"
	_favoriteDelivery "aprian1337/thukul-service/deliveries/favorites"
	_favoriteDb "aprian1337/thukul-service/repository/databases/favorites"

	_walletUsecase "aprian1337/thukul-service/business/wallets"
	_walletDb "aprian1337/thukul-service/repository/databases/wallets"

	_walletHistoryUsecase "aprian1337/thukul-service/business/wallet_histories"
	_walletHistoryDb "aprian1337/thukul-service/repository/databases/wallet_histories"

	_paymentsUsecase "aprian1337/thukul-service/business/payments"
	_paymentDelivery "aprian1337/thukul-service/deliveries/payments"

	_cryptosUsecase "aprian1337/thukul-service/business/cryptos"
	_cryptoDb "aprian1337/thukul-service/repository/databases/cryptos"

	_transactionUsecase "aprian1337/thukul-service/business/transactions"
	_transactionDb "aprian1337/thukul-service/repository/databases/transactions"

	"aprian1337/thukul-service/repository/drivers/mongodb"
	"aprian1337/thukul-service/repository/drivers/postgres"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"time"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service run on DEBUG MODE")
	}
}

func DbMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&_salaryDb.Salaries{},
		&_usersDb.Users{},
		&_pocketDb.Pockets{},
		&_activityDb.Activities{},
		&_coinDb.Coins{},
		&_favoriteDb.Favorites{},
		&_wishlistDb.Wishlists{},
		&_transactionHistoryDb.Transactions{},
		&_walletDb.Wallets{},
		&_walletHistoryDb.WalletHistories{},
	)
	if err != nil {
		panic(err)
	}
}

func main() {
	postgresConfig := postgres.ConfigDb{
		DbHost:     viper.GetString(`databases.postgres.host`),
		DbUser:     viper.GetString(`databases.postgres.user`),
		DbPassword: viper.GetString(`databases.postgres.password`),
		DbName:     viper.GetString(`databases.postgres.dbname`),
		DbPort:     viper.GetString(`databases.postgres.port`),
		DbSslMode:  viper.GetString(`databases.postgres.sslmode`),
		DbTimezone: viper.GetString(`databases.postgres.timezone`),
	}

	mongoConfig := mongodb.ConfigDb{
		DbHost: viper.GetString(`databases.mongodb.host`),
		DbPort: viper.GetString(`databases.mongodb.port`),
	}

	configJWT := middlewares.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	connPostgres := postgresConfig.InitialDb(viper.GetBool(`debug`))

	//MONGO
	logCol := middlewares.InitCollection(struct {
		DbName     string
		Collection string
	}{
		DbName:     viper.GetString(`databases.mongodb.dbname`),
		Collection: viper.GetString(`databases.mongodb.collection.logger`),
	})

	initMongo := mongoConfig.InitDb()
	loggerMiddleware := middlewares.InitConfig(initMongo, logCol, viper.GetDuration(`context.timeout`))

	DbMigrate(connPostgres)
	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	configMarketRepo := _coinmarketRepo.MarketCapAPI{
		BaseUrl:        constants.BaseUrlApiMarketcap,
		EndpointSymbol: constants.EndpointMarketcapSymbol,
		EndpointPrice:  constants.EndpointMarketcapPrice,
		ApiKey:         viper.GetString("thirdparties.coinmarketcap.api_key"),
	}
	coinMarketRepo := _coinmarketRepo.NewMarketCapAPI(configMarketRepo)

	cryptoRepository := _cryptoDb.NewPostgresCryptosRepository(connPostgres)
	cryptoUsecase := _cryptosUsecase.NewCryptoUsecase(cryptoRepository, timeoutContext)

	coinRepository := _coinDb.NewPostgresCoinsRepository(connPostgres)
	coinUsecase := _coinUsecase.NewCoinUsecase(coinRepository, coinMarketRepo, timeoutContext)
	coinDelivery := _coinDelivery.NewCoinsController(coinUsecase)

	walletsHistoryRepository := _walletHistoryDb.NewPostgresWalletHistoriesRepository(connPostgres)
	walletsHistoryUsecase := _walletHistoryUsecase.NewWalletsUsecase(walletsHistoryRepository, timeoutContext)

	walletsRepository := _walletDb.NewPostgresWalletsRepository(connPostgres)
	walletsUsecase := _walletUsecase.NewWalletsUsecase(walletsRepository, walletsHistoryUsecase, timeoutContext)

	transactionsRepository := _transactionDb.NewPostgresTransactionRepository(connPostgres)
	transactionsUsecase := _transactionUsecase.NewTransactionUsecase(transactionsRepository, timeoutContext)

	paymentUsecase := _paymentsUsecase.NewPaymentUsecase(cryptoUsecase, coinUsecase, coinMarketRepo, walletsUsecase, walletsHistoryUsecase, transactionsUsecase, timeoutContext)
	paymentDelivery := _paymentDelivery.NewFavoriteController(paymentUsecase)

	userRepository := _usersDb.NewPostgresUserRepository(connPostgres)
	userUsecase := _usersUsecase.NewUserUsecase(userRepository, walletsUsecase, timeoutContext, &configJWT)
	userDelivery := _usersDelivery.NewUserController(userUsecase)

	salaryRepository := _salaryDb.NewPostgresSalariesRepository(connPostgres)
	salaryUsecase := _salaryUsecase.NewSalaryUsecase(salaryRepository, timeoutContext)
	salaryDelivery := _salaryDelivery.NewSalariesController(salaryUsecase)

	activityRepository := _activityDb.NewPostgresPocketsRepository(connPostgres)
	activityUsecase := _activityUsecase.NewActivityUsecase(activityRepository, timeoutContext)
	activityDelivery := _activityDelivery.NewActivityController(activityUsecase)

	pocketRepository := _pocketDb.NewPostgresPocketsRepository(connPostgres)
	pocketUsecase := _pocketUsecase.NewPocketUsecase(pocketRepository, activityUsecase, timeoutContext)
	pocketDelivery := _pocketDelivery.NewSalariesController(pocketUsecase)

	wishlistRepository := _wishlistDb.NewPostgresWishlistRepository(connPostgres)
	wishlistUsecase := _wishlistUsecase.NewWishlistUsecase(wishlistRepository, userUsecase, timeoutContext)
	wishlistDelivery := _wishlistDelivery.NewSalariesController(wishlistUsecase)

	favoriteRepository := _favoriteDb.NewPostgresFavoritesRepository(connPostgres)
	favoriteUsecase := _favoriteUsecase.NewFavoriteUsecase(favoriteRepository, userUsecase, coinUsecase, timeoutContext)
	favoriteDelivery := _favoriteDelivery.NewFavoriteController(favoriteUsecase)

	routesInit := routes.ControllerList{
		UserController:     *userDelivery,
		SalaryController:   *salaryDelivery,
		PocketController:   *pocketDelivery,
		ActivityController: *activityDelivery,
		CoinController:     *coinDelivery,
		WishlistController: *wishlistDelivery,
		FavoriteController: *favoriteDelivery,
		PaymentController:  *paymentDelivery,
		LoggerMiddleware:   *loggerMiddleware,
		JWTMiddleware:      configJWT.Init(),
	}

	routesInit.RouteUsers(e)
	address := fmt.Sprintf("%v:%v",
		viper.GetString("server.address.host"),
		viper.GetString("server.address.port"),
	)
	err := e.Start(address)
	if err != nil {
		panic(err)
	}
}
