package main

import (
	"aprian1337/thukul-service/app/middlewares"
	"aprian1337/thukul-service/app/routes"
	_usersUsecase "aprian1337/thukul-service/business/users"
	_usersDelivery "aprian1337/thukul-service/deliveries/users"
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
	middlewareConf := middlewares.InitConfig(initMongo, logCol)

	DbMigrate(connPostgres)
	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	configMarketRepo := _coinmarketRepo.MarketCapAPI{
		BaseUrl:        viper.GetString("thirdparties.coinmarketcap.base_url"),
		ApiKey:         viper.GetString("thirdparties.coinmarketcap.api_key"),
		EndpointSymbol: viper.GetString("thirdparties.coinmarketcap.endpoint_symbol"),
	}
	coinMarketRepo := _coinmarketRepo.NewMarketCapAPI(configMarketRepo)

	userRepository := _usersDb.NewPostgresUserRepository(connPostgres)
	userUsecase := _usersUsecase.NewUserUsecase(userRepository, timeoutContext, &configJWT)
	userDelivery := _usersDelivery.NewUserController(userUsecase)

	salaryRepository := _salaryDb.NewPostgresSalariesRepository(connPostgres)
	salaryUsecase := _salaryUsecase.NewSalaryUsecase(salaryRepository, timeoutContext)
	salaryDelivery := _salaryDelivery.NewSalariesController(salaryUsecase)

	activityRepository := _activityDb.NewPostgresPocketsRepository(connPostgres)
	activityUsecase := _activityUsecase.NewActivityUsecase(activityRepository, timeoutContext)
	activityDelivery := _activityDelivery.NewActivityController(activityUsecase)

	pocketRepository := _pocketDb.NewPostgresPocketsRepository(connPostgres)
	pocketUsecase := _pocketUsecase.NewPocketUsecase(pocketRepository, activityRepository, timeoutContext)
	pocketDelivery := _pocketDelivery.NewSalariesController(pocketUsecase)

	coinRepository := _coinDb.NewPostgresCoinsRepository(connPostgres)
	coinUsecase := _coinUsecase.NewCoinUsecase(coinRepository, coinMarketRepo, timeoutContext)
	coinDelivery := _coinDelivery.NewCoinsController(coinUsecase)

	routesInit := routes.ControllerList{
		MiddlewareConfig:   *middlewareConf,
		UserController:     *userDelivery,
		SalaryController:   *salaryDelivery,
		PocketController:   *pocketDelivery,
		ActivityController: *activityDelivery,
		CoinController:     *coinDelivery,
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
