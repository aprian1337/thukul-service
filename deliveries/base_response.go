package deliveries

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type BaseResponse struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	Data interface{}
}

func NewSuccessResponse(c echo.Context, data interface{}) error {
	response := BaseResponse{
		Meta: struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}{
			Status:  http.StatusOK,
			Message: "Success",
		},
		Data: data,
	}
	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{
		Meta: struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}{
			Status:  status,
			Message: err.Error(),
		},
		Data: nil,
	}
	return c.JSON(status, response)
}
