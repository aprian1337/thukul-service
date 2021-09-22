package config

import (
	"aprian1337/thukul-service/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

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
	err := Db.AutoMigrate(&models.Salaries{}, &models.Users{})
	if err != nil {
		return
	}
}
