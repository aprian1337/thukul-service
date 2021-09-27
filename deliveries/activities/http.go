package activities

import (
	"aprian1337/thukul-service/business/activities"
	"aprian1337/thukul-service/deliveries"
	"aprian1337/thukul-service/deliveries/activities/requests"
	"aprian1337/thukul-service/deliveries/activities/responses"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	ActivityController activities.Usecase
}

func NewActivityController(uc activities.Usecase) *Controller {
	return &Controller{
		ActivityController: uc,
	}
}

func (ctrl *Controller) Get(c echo.Context) error {
	ctx := c.Request().Context()
	idPocket := c.Param("idPocket")
	idPocketConv, err := strconv.Atoi(idPocket)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	data, err := ctrl.ActivityController.GetList(ctx, idPocketConv)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromListDomain(data))
}

func (ctrl *Controller) GetById(c echo.Context) error {
	pocketId := c.Param("idPocket")
	convPocketId, err := strconv.Atoi(pocketId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()

	data, err := ctrl.ActivityController.GetById(ctx, convPocketId, convId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomain(data))
}

func (ctrl *Controller) Update(c echo.Context) error {
	id := c.Param("id")
	request := requests.ActivityRequest{}
	err := c.Bind(&request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	idPocket := c.Param("idPocket")
	idPocketConv, err := strconv.Atoi(idPocket)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	convId, err := strconv.Atoi(id)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	data, err := ctrl.ActivityController.Update(ctx, request.ToDomain(), idPocketConv, convId)
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
	idPocket := c.Param("idPocket")
	idPocketConv, err := strconv.Atoi(idPocket)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	err = ctrl.ActivityController.Delete(ctx, convId, idPocketConv)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.ActivitiesResponse{
		ID: convId,
	})
}

func (ctrl *Controller) Create(c echo.Context) error {
	request := requests.ActivityRequest{}
	err := c.Bind(&request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	idPocket := c.Param("idPocket")
	idPocketConv, err := strconv.Atoi(idPocket)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	data, err := ctrl.ActivityController.Create(ctx, request.ToDomain(), idPocketConv)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomain(data))
}
