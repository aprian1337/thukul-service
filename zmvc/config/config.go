package config

import (
	"aprian1337/thukul-service/zmvc/models/salaries"
	"aprian1337/thukul-service/zmvc/models/users"
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetClientMongo() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func InitDB() {
	dsn := "host=localhost user=aprian1337 password=root123 dbname=thukul port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	Migration()
}

func Migration() {
	fmt.Println("Migration..")
	err := Db.AutoMigrate(&salaries.Db{}, &users.Db{})
	if err != nil {
		return
	}
}
