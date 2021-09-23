package controllers

import (
	"aprian1337/thukul-service/zmvc/config"
	"aprian1337/thukul-service/zmvc/models/salaries"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetSalariesController(c echo.Context) error {
	var getSalaries []salaries.Db

	err := config.Db.Find(&getSalaries).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    getSalaries,
	})
}

func CreateSalariesController(context echo.Context) error {
	salary := salaries.Db{}
	err := context.Bind(&salary)
	if err != nil {
		panic(err.Error())
		return err
	}
	errDb := config.Db.Create(&salary).Error
	if errDb != nil {
		return context.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return context.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    salary,
	})
}
