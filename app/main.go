package main

import (
	"aprian1337/thukul-service/app/middlewares"
	"aprian1337/thukul-service/app/routes"
	_smtpUsecase "aprian1337/thukul-service/business/smtp"
	_usersUsecase "aprian1337/thukul-service/business/users"
	_usersDelivery "aprian1337/thukul-service/deliveries/users"
	"aprian1337/thukul-service/helpers/constants"
	postgres2 "aprian1337/thukul-service/repository/databases/postgres"
	_activityDb "aprian1337/thukul-service/repository/databases/records"
	_coinmarketRepo "aprian1337/thukul-service/repository/thirdparties/coinmarket"

	_activityUsecase "aprian1337/thukul-service/business/activities"
	_coinUsecase "aprian1337/thukul-service/business/coins"
	_favoriteUsecase "aprian1337/thukul-service/business/favorites"
	_paymentsUsecase "aprian1337/thukul-service/business/payments"
	_pocketUsecase "aprian1337/thukul-service/business/pockets"
	_salaryUsecase "aprian1337/thukul-service/business/salaries"
	_walletHistoryUsecase "aprian1337/thukul-service/business/wallet_histories"
	_walletUsecase "aprian1337/thukul-service/business/wallets"
	_wishlistUsecase "aprian1337/thukul-service/business/wishlists"
	_activityDelivery "aprian1337/thukul-service/deliveries/activities"
	_coinDelivery "aprian1337/thukul-service/deliveries/coins"
	_favoriteDelivery "aprian1337/thukul-service/deliveries/favorites"
	_paymentDelivery "aprian1337/thukul-service/deliveries/payments"
	_pocketDelivery "aprian1337/thukul-service/deliveries/pockets"
	_salaryDelivery "aprian1337/thukul-service/deliveries/salaries"
	_wishlistDelivery "aprian1337/thukul-service/deliveries/wishlists"

	_cryptosUsecase "aprian1337/thukul-service/business/cryptos"
	_transactionUsecase "aprian1337/thukul-service/business/transactions"
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
		&_activityDb.Salaries{},
		&_activityDb.Users{},
		&_activityDb.Pockets{},
		&_activityDb.Activities{},
		&_activityDb.Coins{},
		&_activityDb.Favorites{},
		&_activityDb.Wishlists{},
		&_activityDb.Transactions{},
		&_activityDb.Wallets{},
		&_activityDb.WalletHistories{},
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

	smtpUsecase := _smtpUsecase.NewSmtpUsecase(
		viper.GetString(`smtp.host`),
		viper.GetInt(`smtp.port`),
		viper.GetString(`smtp.sender_name`),
		viper.GetString(`smtp.email`),
		viper.GetString(`smtp.password`),
	)

	coinMarketRepo := _coinmarketRepo.NewMarketCapAPI(configMarketRepo)

	cryptoRepository := postgres2.NewPostgresCryptosRepository(connPostgres)
	cryptoUsecase := _cryptosUsecase.NewCryptoUsecase(cryptoRepository, timeoutContext)

	coinRepository := postgres2.NewPostgresCoinsRepository(connPostgres)
	coinUsecase := _coinUsecase.NewCoinUsecase(coinRepository, coinMarketRepo, timeoutContext)
	coinDelivery := _coinDelivery.NewCoinsController(coinUsecase)

	walletsHistoryRepository := postgres2.NewPostgresWalletHistoriesRepository(connPostgres)
	walletsHistoryUsecase := _walletHistoryUsecase.NewWalletsUsecase(walletsHistoryRepository, timeoutContext)

	walletsRepository := postgres2.NewPostgresWalletsRepository(connPostgres)
	walletsUsecase := _walletUsecase.NewWalletsUsecase(walletsRepository, walletsHistoryUsecase, timeoutContext)

	transactionsRepository := postgres2.NewPostgresTransactionRepository(connPostgres)
	transactionsUsecase := _transactionUsecase.NewTransactionUsecase(transactionsRepository, timeoutContext)

	userRepository := postgres2.NewPostgresUserRepository(connPostgres)
	userUsecase := _usersUsecase.NewUserUsecase(userRepository, walletsUsecase, timeoutContext, &configJWT)
	userDelivery := _usersDelivery.NewUserController(userUsecase)

	paymentUsecase := _paymentsUsecase.NewPaymentUsecase(userUsecase, smtpUsecase, cryptoUsecase, coinUsecase, coinMarketRepo, walletsUsecase, walletsHistoryUsecase, transactionsUsecase, viper.GetString(`encrypt.keystring`), viper.GetString(`encrypt.additional`), viper.GetString("server.address.host"), timeoutContext)
	paymentDelivery := _paymentDelivery.NewFavoriteController(paymentUsecase)

	salaryRepository := postgres2.NewPostgresSalariesRepository(connPostgres)
	salaryUsecase := _salaryUsecase.NewSalaryUsecase(salaryRepository, timeoutContext)
	salaryDelivery := _salaryDelivery.NewSalariesController(salaryUsecase)

	activityRepository := postgres2.NewPostgresActivitiesRepository(connPostgres)
	activityUsecase := _activityUsecase.NewActivityUsecase(activityRepository, timeoutContext)
	activityDelivery := _activityDelivery.NewActivityController(activityUsecase)

	pocketRepository := postgres2.NewPostgresPocketsRepository(connPostgres)
	pocketUsecase := _pocketUsecase.NewPocketUsecase(pocketRepository, activityUsecase, timeoutContext)
	pocketDelivery := _pocketDelivery.NewSalariesController(pocketUsecase)

	wishlistRepository := postgres2.NewPostgresWishlistRepository(connPostgres)
	wishlistUsecase := _wishlistUsecase.NewWishlistUsecase(wishlistRepository, userUsecase, timeoutContext)
	wishlistDelivery := _wishlistDelivery.NewSalariesController(wishlistUsecase)

	favoriteRepository := postgres2.NewPostgresFavoritesRepository(connPostgres)
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
