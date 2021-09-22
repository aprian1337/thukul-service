package controllers

import (
	"aprian1337/thukul-service/config"
	"aprian1337/thukul-service/helpers"
	"aprian1337/thukul-service/models/responses"
	"aprian1337/thukul-service/models/users"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func GetUsersController(ctx echo.Context) error {
	var listUsers []users.Db

	err := config.Db.Find(&listUsers).Error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    listUsers,
	})
}

func CreateUsersController(ctx echo.Context) error {
	newUser := users.Request{}
	err := ctx.Bind(&newUser)
	if err != nil {
		panic(err.Error())
		return err
	}
	newUser.Password = helpers.HashWithBcrypt(newUser.Password)
	var userDb users.Db
	birthday, errTime := time.Parse("2006-01-02", newUser.Birthday)
	if errTime != nil {
		return err
	}
	userDb = users.Db{
		Name:     newUser.Name,
		SalaryId: newUser.SalaryId,
		Birthday: birthday,
		Address:  newUser.Address,
		Password: newUser.Password,
		Company:  newUser.Company,
		Gender:   newUser.Gender,
		IsAdmin:  newUser.IsAdmin,
		IsValid:  newUser.IsValid,
		Phone:    newUser.Phone,
		Email:    newUser.Email,
	}
	errDb := config.Db.Save(&userDb).Error
	if errDb != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    userDb,
	})
}

func LoginUsersController(ctx echo.Context) error {
	var userLogin users.UserLogin
	err := ctx.Bind(&userLogin)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
		})
	}
	if userLogin.Email == "" || userLogin.Password == "" {
		return ctx.JSON(http.StatusBadRequest, responses.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Username or password must be declare",
			Data:    nil,
		})
	}
	var userDb users.Db
	result := config.Db.First(&userDb, "email = ? ", userLogin.Email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusForbidden, responses.BaseResponse{
				Code:    http.StatusForbidden,
				Message: "User or password is not valid",
				Data:    nil,
			})
		} else {
			return ctx.JSON(http.StatusInternalServerError, responses.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Server not respond",
				Data:    nil,
			})
		}
	}
	if helpers.CompareHashWithBcrypt(userDb.Password, userLogin.Password) == false {
		return ctx.JSON(http.StatusForbidden, responses.BaseResponse{
			Code:    http.StatusForbidden,
			Message: "User or password is not valid",
			Data:    nil,
		})
	}
	userResponse := users.Response{
		Id:       userDb.ID,
		SalaryId: userDb.SalaryId,
		Name:     userDb.Name,
		IsAdmin:  userDb.IsAdmin,
		Email:    userDb.Email,
		Phone:    userDb.Phone,
		Gender:   userDb.Gender,
		Birthday: userDb.Birthday.Format("2006-01-02"),
		Address:  userDb.Address,
		Company:  userDb.Company,
		IsValid:  userDb.IsValid,
	}
	return ctx.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    userResponse,
	})
}
