package main

import (
	"aprian1337/thukul-service/app/middlewares"
	"aprian1337/thukul-service/app/routes"
	_coinsDeliveryGql "aprian1337/thukul-service/deliveries/coins/graphql"
	_paymentsDeliveryGql "aprian1337/thukul-service/deliveries/payments/graphql"
	_salaryDeliveryGql "aprian1337/thukul-service/deliveries/salaries/graphql"
	_usersDeliveryGql "aprian1337/thukul-service/deliveries/users/graphql"
	"aprian1337/thukul-service/gql"
	"aprian1337/thukul-service/helpers/constants"
	"aprian1337/thukul-service/repository/drivers/mongodb"
	"aprian1337/thukul-service/repository/drivers/postgres"
	"context"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"

	postgresRepo "aprian1337/thukul-service/repository/databases/postgres"

	_smtpUsecase "aprian1337/thukul-service/business/smtp"
	_coinmarketRepo "aprian1337/thukul-service/repository/thirdparties/coinmarket"

	_usersUsecase "aprian1337/thukul-service/business/users"
	_usersDelivery "aprian1337/thukul-service/deliveries/users"

	_activityUsecase "aprian1337/thukul-service/business/activities"
	_activityDb "aprian1337/thukul-service/repository/databases/records"

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
	_cryptosDelivery "aprian1337/thukul-service/deliveries/cryptos"
)

func init() {
	viper.SetConfigFile(`config.json`)
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
		&_activityDb.Cryptos{},
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
		Cluster:  viper.GetString(`databases.mongodb.cluster`),
		Username: viper.GetString(`databases.mongodb.username`),
		Password: viper.GetString(`databases.mongodb.password`),
		IsLocal:  viper.GetBool(`databases.mongodb.is_local`),
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

	cryptoRepository := postgresRepo.NewPostgresCryptosRepository(connPostgres)
	cryptoUsecase := _cryptosUsecase.NewCryptoUsecase(cryptoRepository, timeoutContext)
	cryptoDelivery := _cryptosDelivery.NewController(cryptoUsecase)

	//HTTP
	coinRepository := postgresRepo.NewPostgresCoinsRepository(connPostgres)
	coinUsecase := _coinUsecase.NewCoinUsecase(coinRepository, coinMarketRepo, timeoutContext)
	coinDelivery := _coinDelivery.NewCoinsController(coinUsecase)
	//GQL
	coinsResolver := _coinsDeliveryGql.NewCoinsResolver(*coinUsecase)
	coinsSchema := _coinsDeliveryGql.NewSchema(*coinsResolver)

	walletsHistoryRepository := postgresRepo.NewPostgresWalletHistoriesRepository(connPostgres)
	walletsHistoryUsecase := _walletHistoryUsecase.NewWalletsUsecase(walletsHistoryRepository, timeoutContext)

	walletsRepository := postgresRepo.NewPostgresWalletsRepository(connPostgres)
	walletsUsecase := _walletUsecase.NewWalletsUsecase(walletsRepository, walletsHistoryUsecase, timeoutContext)

	transactionsRepository := postgresRepo.NewPostgresTransactionRepository(connPostgres)
	transactionsUsecase := _transactionUsecase.NewTransactionUsecase(transactionsRepository, timeoutContext)

	//HTTP
	userRepository := postgresRepo.NewPostgresUserRepository(connPostgres)
	userUsecase := _usersUsecase.NewUserUsecase(userRepository, walletsUsecase, timeoutContext, &configJWT)
	userDelivery := _usersDelivery.NewUserController(userUsecase)
	//GQL
	usersResolver := _usersDeliveryGql.NewUsersResolver(*userUsecase)
	usersSchema := _usersDeliveryGql.NewSchema(*usersResolver)

	//HTTP
	paymentUsecase := _paymentsUsecase.NewPaymentUsecase(userUsecase, smtpUsecase, cryptoUsecase, coinUsecase, coinMarketRepo, walletsUsecase, walletsHistoryUsecase, transactionsUsecase, viper.GetString(`encrypt.keystring`), viper.GetString(`encrypt.additional`), viper.GetString("smtp.server"), viper.GetString("server.address.port"), timeoutContext)
	paymentDelivery := _paymentDelivery.NewFavoriteController(paymentUsecase)
	//GQL
	paymentResolver := _paymentsDeliveryGql.NewPaymentsResolver(*paymentUsecase)
	paymentSchema := _paymentsDeliveryGql.NewSchema(*paymentResolver)

	//HTTP
	salaryRepository := postgresRepo.NewPostgresSalariesRepository(connPostgres)
	salaryUsecase := _salaryUsecase.NewSalaryUsecase(salaryRepository, timeoutContext)
	salaryDelivery := _salaryDelivery.NewSalariesController(salaryUsecase)
	//GQL
	salaryResolver := _salaryDeliveryGql.NewSalaryResolver(*salaryUsecase)
	salarySchema := _salaryDeliveryGql.NewSchema(*salaryResolver)

	activityRepository := postgresRepo.NewPostgresActivitiesRepository(connPostgres)
	activityUsecase := _activityUsecase.NewActivityUsecase(activityRepository, timeoutContext)
	activityDelivery := _activityDelivery.NewActivityController(activityUsecase)

	pocketRepository := postgresRepo.NewPostgresPocketsRepository(connPostgres)
	pocketUsecase := _pocketUsecase.NewPocketUsecase(pocketRepository, activityUsecase, timeoutContext)
	pocketDelivery := _pocketDelivery.NewSalariesController(pocketUsecase)

	wishlistRepository := postgresRepo.NewPostgresWishlistRepository(connPostgres)
	wishlistUsecase := _wishlistUsecase.NewWishlistUsecase(wishlistRepository, userUsecase, timeoutContext)
	wishlistDelivery := _wishlistDelivery.NewSalariesController(wishlistUsecase)

	favoriteRepository := postgresRepo.NewPostgresFavoritesRepository(connPostgres)
	favoriteUsecase := _favoriteUsecase.NewFavoriteUsecase(favoriteRepository, userUsecase, coinUsecase, timeoutContext)
	favoriteDelivery := _favoriteDelivery.NewFavoriteController(favoriteUsecase)

	schema := gql.Schema{
		SalarySchema:  salarySchema,
		CoinsSchema:   coinsSchema,
		PaymentSchema: paymentSchema,
		UsersSchema:   usersSchema,
	}
	gqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    schema.Query(),
		Mutation: schema.Mutation(),
	})

	gqlHandler := handler.New(&handler.Config{
		Schema:     &gqlSchema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
		RootObjectFn: func(ctx context.Context, r *http.Request) map[string]interface{} {
			return map[string]interface{}{
				"token": r.Header.Get("token"),
			}
		},
	})

	routesInit := routes.ControllerList{
		UserController:     *userDelivery,
		SalaryController:   *salaryDelivery,
		CryptoController:   *cryptoDelivery,
		PocketController:   *pocketDelivery,
		ActivityController: *activityDelivery,
		CoinController:     *coinDelivery,
		WishlistController: *wishlistDelivery,
		FavoriteController: *favoriteDelivery,
		PaymentController:  *paymentDelivery,
		LoggerMiddleware:   *loggerMiddleware,
		JWTMiddleware:      configJWT.Init(),
		GqlHandler:         gqlHandler,
	}

	routesInit.Route(e)
	address := fmt.Sprintf("%v:%v",
		viper.GetString("server.address.host"),
		viper.GetString("server.address.port"),
	)
	err = e.Start(address)
	if err != nil {
		panic(err)
	}
}
