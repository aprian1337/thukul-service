package favorites

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/favorites"
	"aprian1337/thukul-service/deliveries"
	"aprian1337/thukul-service/deliveries/favorites/requests"
	"aprian1337/thukul-service/deliveries/favorites/responses"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	favoriteUsecase favorites.Usecase
}

func NewFavoriteController(uc favorites.Usecase) *Controller {
	return &Controller{
		favoriteUsecase: uc,
	}
}

func (ctrl *Controller) Get(c echo.Context) error {
	ctxNative := c.Request().Context()
	userId := c.Param("userId")
	convId, err := strconv.Atoi(userId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, businesses.ErrBadRequest)
	}
	data, err := ctrl.favoriteUsecase.GetList(ctxNative, convId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	fmt.Println(data)
	return deliveries.NewSuccessResponse(c, responses.FromListDomain(data))
}

func (ctrl *Controller) GetById(c echo.Context) error {
	ctxNative := c.Request().Context()
	userId := c.Param("userId")
	convId, err := strconv.Atoi(userId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, businesses.ErrBadRequest)
	}
	favId := c.Param("favId")
	convFavId, err := strconv.Atoi(favId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, businesses.ErrBadRequest)
	}
	data, err := ctrl.favoriteUsecase.GetById(ctxNative, convId, convFavId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomain(data))
}

func (ctrl *Controller) Create(c echo.Context) error {
	request := requests.FavoriteRequest{}
	err := c.Bind(&request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	userId := c.Param("userId")
	convId, err := strconv.Atoi(userId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, businesses.ErrBadRequest)
	}
	ctxNative := c.Request().Context()
	var data favorites.Domain
	data, err = ctrl.favoriteUsecase.Create(ctxNative, request.ToDomain(), convId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomain(data))
}

func (ctrl *Controller) Destroy(c echo.Context) error {
	userId := c.Param("userId")
	convId, err := strconv.Atoi(userId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, businesses.ErrBadRequest)
	}
	favId := c.Param("favId")
	convFavId, err := strconv.Atoi(favId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, businesses.ErrBadRequest)
	}
	ctx := c.Request().Context()
	err = ctrl.favoriteUsecase.Delete(ctx, convId, convFavId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, nil)
}
