package middlewares

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MiddlewareConfig struct {
	Mongo *mongo.Client
	Logs  *LogCollection
}

type LogCollection struct {
	DbName     string
	Collection string
}

type Logger struct {
	Uri      string
	Method   string
	Status   int
	UserIp   string
	HostIp   string
	Time     time.Time
	Response string
}

func InitConfig(db *mongo.Client, logs *LogCollection) *MiddlewareConfig {
	return &MiddlewareConfig{
		Mongo: db,
		Logs:  logs,
	}
}

func InitCollection(logs LogCollection) *LogCollection {
	return &LogCollection{
		DbName:     logs.DbName,
		Collection: logs.Collection,
	}
}

func (mc *MiddlewareConfig) Start(e *echo.Echo) {
	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: middleware.DefaultSkipper,
		Handler: func(e echo.Context, req []byte, resp []byte) {
			collection := mc.Mongo.Database(mc.Logs.DbName).Collection(mc.Logs.Collection)
			logs := Logger{
				Uri:      e.Request().RequestURI,
				Method:   e.Request().Method,
				Status:   e.Response().Status,
				UserIp:   e.RealIP(),
				HostIp:   e.Request().Host,
				Time:     time.Now().Local(),
				Response: string(resp),
			}
			ctx, cancel := context.WithTimeout(context.Background(), 22*time.Second)
			defer cancel()
			_, err := collection.InsertOne(ctx, logs)
			if err != nil {
				panic(err)
			}
		},
	}))
	e.Pre(middleware.RemoveTrailingSlash())
}
