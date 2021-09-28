package pockets

import (
	"aprian1337/thukul-service/business/pockets"
	"aprian1337/thukul-service/deliveries"
	"aprian1337/thukul-service/deliveries/pockets/requests"
	"aprian1337/thukul-service/deliveries/pockets/responses"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	PocketsUsecase pockets.Usecase
}

func NewSalariesController(uc pockets.Usecase) *Controller {
	return &Controller{
		PocketsUsecase: uc,
	}
}

func (ctrl *Controller) Get(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Param("userId")
	data, err := ctrl.PocketsUsecase.GetList(ctx, userId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromListDomain(data))
}

func (ctrl *Controller) GetById(c echo.Context) error {
	userId := c.Param("userId")
	convUserId, err := strconv.Atoi(userId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	pocketId := c.Param("pocketId")
	convPocketId, err := strconv.Atoi(pocketId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()

	data, err := ctrl.PocketsUsecase.GetById(ctx, convUserId, convPocketId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomain(data))
}

func (ctrl *Controller) Update(c echo.Context) error {
	userId := c.Param("userId")
	convUserId, err := strconv.Atoi(userId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	pocketId := c.Param("pocketId")
	convPocketId, err := strconv.Atoi(pocketId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	request := requests.PocketsRequest{}
	err = c.Bind(&request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()

	data, err := ctrl.PocketsUsecase.Update(ctx, request.ToDomain(), convUserId, convPocketId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomain(data))
}

func (ctrl *Controller) Destroy(c echo.Context) error {
	userId := c.Param("userId")
	convUserId, err := strconv.Atoi(userId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	pocketId := c.Param("pocketId")
	convPocketId, err := strconv.Atoi(pocketId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	err = ctrl.PocketsUsecase.Delete(ctx, convUserId, convPocketId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.PocketsResponse{
		ID: convUserId,
	})
}

func (ctrl *Controller) Create(c echo.Context) error {
	request := requests.PocketsRequest{}
	err := c.Bind(&request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	userId := c.Param("userId")
	convUserId, err := strconv.Atoi(userId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	request.UserId = convUserId

	ctx := c.Request().Context()
	data, err := ctrl.PocketsUsecase.Create(ctx, request.ToDomain())
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomain(data))
}

func (ctrl *Controller) Total(c echo.Context) error {
	kind := c.QueryParam("type")
	userId := c.Param("userId")
	pocketId := c.Param("pocketId")
	convUserId, err := strconv.Atoi(userId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	convPocketId, err := strconv.Atoi(pocketId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	data, err := ctrl.PocketsUsecase.GetTotalByActivities(ctx, convUserId, convPocketId, kind)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, data)
}
