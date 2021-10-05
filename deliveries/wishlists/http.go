package wishlists

import (
	"aprian1337/thukul-service/business/wishlists"
	"aprian1337/thukul-service/deliveries"
	"aprian1337/thukul-service/deliveries/wishlists/requests"
	"aprian1337/thukul-service/deliveries/wishlists/responses"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	WishlistUsecase wishlists.Usecase
}

func NewSalariesController(uc wishlists.Usecase) *Controller {
	return &Controller{
		WishlistUsecase: uc,
	}
}

func (ctrl *Controller) Get(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Param("userId")
	convId, err := strconv.Atoi(userId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	data, err := ctrl.WishlistUsecase.GetList(ctx, convId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromListDomain(data))
}

func (ctrl *Controller) GetById(c echo.Context) error {
	userId := c.Param("userId")
	wishlistId := c.Param("wishlistId")
	convUserId, err := strconv.Atoi(userId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	convWishlistId, err := strconv.Atoi(wishlistId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()

	data, err := ctrl.WishlistUsecase.GetById(ctx, convUserId, convWishlistId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomain(data))
}

func (ctrl *Controller) Update(c echo.Context) error {
	userId := c.Param("userId")
	wishlistId := c.Param("wishlistId")
	convUserId, err := strconv.Atoi(userId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	convWishlistId, err := strconv.Atoi(wishlistId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	request := requests.WishlistsRequest{}
	err = c.Bind(&request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	data, err := ctrl.WishlistUsecase.Update(ctx, request.ToDomain(), convUserId, convWishlistId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomain(data))
}

func (ctrl *Controller) Destroy(c echo.Context) error {
	userId := c.Param("userId")
	wishlistId := c.Param("wishlistId")
	convUserId, err := strconv.Atoi(userId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	convWishlistId, err := strconv.Atoi(wishlistId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = ctrl.WishlistUsecase.Delete(ctx, convUserId, convWishlistId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.WishlistsResponse{
		ID: convWishlistId,
	})
}

func (ctrl *Controller) Create(c echo.Context) error {
	request := requests.WishlistsRequest{}
	userId := c.Param("userId")
	convUserId, err := strconv.Atoi(userId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	err = c.Bind(&request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	data, err := ctrl.WishlistUsecase.Create(ctx, request.ToDomain(), convUserId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomain(data))
}
