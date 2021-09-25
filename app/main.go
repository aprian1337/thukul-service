package main

import (
	"aprian1337/thukul-service/app/middlewares"
	"aprian1337/thukul-service/app/routes"

	_usersUsecase "aprian1337/thukul-service/business/users"
	_usersDelivery "aprian1337/thukul-service/deliveries/users"
	_usersDb "aprian1337/thukul-service/repository/databases/users"

	_salaryUsecase "aprian1337/thukul-service/business/salaries"
	_salaryDelivery "aprian1337/thukul-service/deliveries/salaries"
	_salaryDb "aprian1337/thukul-service/repository/databases/salaries"

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
		&_usersDb.Users{},
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
		DbHost: viper.GetString(`databases.mongodbb.host`),
		DbPort: viper.GetString(`databases.mongodbb.port`),
	}

	configJWT := middlewares.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	connPostgres := postgresConfig.InitialDb(viper.GetBool(`debug`))
	mongoConfig.InitDb()

	DbMigrate(connPostgres)
	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := _usersDb.NewPostgresUserRepository(connPostgres)
	userUsecase := _usersUsecase.NewUserUsecase(userRepository, timeoutContext, &configJWT)
	userDelivery := _usersDelivery.NewUserController(userUsecase)

	salaryRepository := _salaryDb.NewPostgresSalariesRepository(connPostgres)
	salaryUsecase := _salaryUsecase.NewSalaryUsecase(salaryRepository, timeoutContext)
	salaryDelivery := _salaryDelivery.NewSalariesController(salaryUsecase)

	routesInit := routes.ControllerList{
		UserController:   *userDelivery,
		SalaryController: *salaryDelivery,
		JWTMiddleware:    configJWT.Init(),
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
