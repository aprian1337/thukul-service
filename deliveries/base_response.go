package deliveries

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type BaseResponse struct {
	Status struct {
		Type     string   `json:"status"`
		Code     int      `json:"code"`
		Message  string   `json:"message"`
		Messages []string `json:"messages"`
	}
	Data interface{}
}

func NewSuccessResponse(c echo.Context, data interface{}) error {
	response := BaseResponse{}
	response.Status.Code = http.StatusOK
	response.Status.Message = "Success"
	response.Status.Type = "success"
	response.Data = data
	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Status.Code = status
	response.Status.Type = "error"
	response.Status.Message = "Something not working properly :("
	response.Status.Messages = []string{err.Error()}
	return c.JSON(status, response)
}
