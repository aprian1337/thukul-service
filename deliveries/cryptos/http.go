package cryptos

import (
	"aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/cryptos"
	"aprian1337/thukul-service/deliveries"
	"aprian1337/thukul-service/deliveries/cryptos/responses"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	CryptoUsecase cryptos.Usecase
}

func NewController(uc cryptos.Usecase) *Controller {
	return &Controller{
		CryptoUsecase: uc,
	}
}

func (ctrl *Controller) GetByUser(c echo.Context) error {
	ctxNative := c.Request().Context()
	userId := c.Param("userId")
	convId, err := strconv.Atoi(userId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, businesses.ErrBadRequest)
	}
	data, err := ctrl.CryptoUsecase.CryptosGetByUser(ctxNative, convId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomainCryptoList(data))
}

func (ctrl *Controller) GetDetail(c echo.Context) error {
	ctxNative := c.Request().Context()
	userId := c.Param("userId")
	convUserId, err := strconv.Atoi(userId)
	cryptoId := c.Param("cryptoId")
	convCryptoId, err := strconv.Atoi(cryptoId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, businesses.ErrBadRequest)
	}
	data, err := ctrl.CryptoUsecase.CryptosGetDetail(ctxNative, convUserId, convCryptoId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomainCrypto(data))
}
