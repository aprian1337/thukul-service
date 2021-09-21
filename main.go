package main

import (
	"aprian1337/thukul-service/config"
	"aprian1337/thukul-service/routes"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	config.InitDB()
	e := routes.V1()
	if err := e.Start(":1231"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// Route
