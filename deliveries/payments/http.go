package payments

import (
	"aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/payments"
	"aprian1337/thukul-service/deliveries"
	"aprian1337/thukul-service/deliveries/payments/requests"
	"aprian1337/thukul-service/deliveries/payments/responses"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
	response := responses.FromDomainWallets(pay)
	return deliveries.NewSuccessResponse(c, responses.TopUpResponse{
		Message: "top up has been success",
		Data:    &response,
	})
}

func (ctrl *Controller) Buy(c echo.Context) error {
	ctxNative := c.Request().Context()
	var data requests.PaymentRequest
	err := c.Bind(&data)
	userId := c.Param("userId")
	convId, err := strconv.Atoi(userId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, businesses.ErrBadRequest)
	}
	data.UserId = convId
	if err != nil {
		return err
	}
	err = ctrl.PaymentUsecase.BuyCoin(ctxNative, data.ToDomain())
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.BuySaleResponse{
		Status:  "success",
		Message: "check your email for confirm the purchase",
	})
}

func (ctrl *Controller) Sell(c echo.Context) error {
	ctxNative := c.Request().Context()
	var data requests.PaymentRequest
	err := c.Bind(&data)
	userId := c.Param("userId")
	convId, err := strconv.Atoi(userId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, businesses.ErrBadRequest)
	}
	data.UserId = convId
	if err != nil {
		return err
	}
	err = ctrl.PaymentUsecase.SellCoin(ctxNative, data.ToDomain())
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.BuySaleResponse{
		Status:  "success",
		Message: "check your email for sales confirmation",
	})
}

func (ctrl *Controller) Confirm(c echo.Context) error {
	ctxNative := c.Request().Context()
	var data requests.PaymentRequest
	err := c.Bind(&data)
	if err != nil {
		return err
	}
	uuidEncode := c.Param("uuid")
	encrypt := c.Param("encrypt")
	pay, err := ctrl.PaymentUsecase.Confirm(ctxNative, uuidEncode, encrypt)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.ConfirmResponse{
		Message:      "success",
		PaymentTotal: pay.NominalTransaction,
		PaymentType:  pay.Kind,
		WalletSaldo:  pay.Total,
	})
}
