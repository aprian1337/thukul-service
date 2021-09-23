package main

import (
	"aprian1337/thukul-service/zmvc/config"
	"aprian1337/thukul-service/zmvc/routes"
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
)

func main() {
	config.InitDB()
	mongoClient := config.GetClientMongo()
	err := mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	} else {
		fmt.Println("MONGO CONNECTED!")
	}
	e := routes.V1()
	if err := e.Start(":1231"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// Route
