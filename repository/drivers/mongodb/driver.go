package mongodb

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ConfigDb struct {
	Cluster  string
	Username string
	Password string
	IsLocal  bool
}

func (config *ConfigDb) InitDb() *mongo.Client {
	var uri string

	if config.IsLocal {
		uri = "mongodb://localhost:27017"
	} else {
		uri = fmt.Sprintf("mongodb+srv://%v:%v@%v",
			config.Username,
			config.Password,
			config.Cluster)
	}

	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}
