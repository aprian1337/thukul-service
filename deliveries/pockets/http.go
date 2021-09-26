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
	id := c.QueryParam("id")
	data, err := ctrl.PocketsUsecase.GetList(ctx, id)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromListDomain(data))
}

func (ctrl *Controller) GetById(c echo.Context) error {
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()

	data, err := ctrl.PocketsUsecase.GetById(ctx, convId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomain(data))
}

func (ctrl *Controller) Update(c echo.Context) error {
	id := c.Param("id")
	request := requests.PocketsRequest{}
	err := c.Bind(&request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	convId, err := strconv.Atoi(id)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	data, err := ctrl.PocketsUsecase.Update(ctx, convId, request.ToDomain())
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomain(data))
}

func (ctrl *Controller) Destroy(c echo.Context) error {
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = ctrl.PocketsUsecase.Delete(ctx, convId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.PocketsResponse{
		ID: convId,
	})
}

func (ctrl *Controller) Create(c echo.Context) error {
	request := requests.PocketsRequest{}
	err := c.Bind(&request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	data, err := ctrl.PocketsUsecase.Create(ctx, request.ToDomain())
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomain(data))
}
