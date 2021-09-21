package controllers

import (
	"aprian1337/thukul-service/config"
	"aprian1337/thukul-service/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetSalariesController(c echo.Context) error {
	var salaries []models.Salaries

	err := config.Db.Find(&salaries).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    salaries,
	})
}

func CreateSalariesController(context echo.Context) error {
	salaries := models.Salaries{}
	err := context.Bind(&salaries)
	if err != nil {
		panic(err.Error())
		return err
	}
	errDb := config.Db.Save(&salaries).Error
	if errDb != nil {
		return context.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return context.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    salaries,
	})
}
