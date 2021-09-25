package salaries

import (
	"aprian1337/thukul-service/business/salaries"
	"aprian1337/thukul-service/deliveries"
	"aprian1337/thukul-service/deliveries/salaries/requests"
	"aprian1337/thukul-service/deliveries/salaries/responses"
	"aprian1337/thukul-service/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	SalaryUsecase salaries.Usecase
}

func NewSalariesController(uc salaries.Usecase) *Controller {
	return &Controller{
		SalaryUsecase: uc,
	}
}

func (ctrl *Controller) GetSalariesController(c echo.Context) error {
	ctx := c.Request().Context()
	search := c.QueryParam("q")
	data, err := ctrl.SalaryUsecase.GetList(ctx, search)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromListDomain(data))
}

func (ctrl *Controller) GetSalaryByIdController(c echo.Context) error {
	id := c.Param("id")
	convInt, err := strconv.Atoi(id)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	ctx := c.Request().Context()

	data, err := ctrl.SalaryUsecase.GetById(ctx, uint(convInt))
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomain(data))
}

func (ctrl *Controller) UpdateSalaryController(c echo.Context) error {
	request := requests.SalariesRequest{}
	err := c.Bind(&request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	data, err := ctrl.SalaryUsecase.Update(ctx, request.ToDomain())
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomain(data))
}

func (ctrl *Controller) DestroySalaryController(c echo.Context) error {
	id := c.Param("id")
	idUint, err := helpers.StringToUint(id)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = ctrl.SalaryUsecase.Delete(ctx, idUint)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.SalariesResponse{
		ID: idUint,
	})
}

func (ctrl *Controller) CreateSalaryController(c echo.Context) error {
	request := requests.SalariesRequest{}
	err := c.Bind(&request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	data, err := ctrl.SalaryUsecase.Create(ctx, request.ToDomain())
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomain(data))
}
