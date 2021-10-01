package payments

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/payments"
	"aprian1337/thukul-service/deliveries"
	"aprian1337/thukul-service/deliveries/payments/requests"
	"aprian1337/thukul-service/deliveries/payments/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	PaymentUsecase payments.Usecase
}

func NewFavoriteController(uc payments.Usecase) *Controller {
	return &Controller{
		PaymentUsecase: uc,
	}
}

func (ctrl *Controller) TopUp(c echo.Context) error {
	ctxNative := c.Request().Context()
	var data requests.PaymentRequest
	err := c.Bind(&data)
	if err != nil {
		return err
	}
	pay, err := ctrl.PaymentUsecase.TopUp(ctxNative, data.ToDomain())
	if err != nil {
		if err == businesses.ErrBadRequest {
			return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
		} else if err == businesses.ErrUserIdNotFound {
			return deliveries.NewErrorResponse(c, http.StatusForbidden, err)
		}
	}
	return deliveries.NewSuccessResponse(c, responses.TopUpResponse{
		Message: "Top up has been success",
		Data:    pay,
	})
}

func (ctrl *Controller) Buy(c echo.Context) error {
	ctxNative := c.Request().Context()
	var data requests.PaymentRequest
	err := c.Bind(&data)
	if err != nil {
		return err
	}
	err = ctrl.PaymentUsecase.BuyCoin(ctxNative, data.ToDomain())
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
		//if err == businesses.ErrBadRequest {
		//	return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
		//} else if err == businesses.ErrUserIdNotFound {
		//	return deliveries.NewErrorResponse(c, http.StatusForbidden, err)
		//}
	}
	return deliveries.NewSuccessResponse(c, responses.TopUpResponse{
		Message: "Buy coin success",
		Data:    nil,
	})
}
