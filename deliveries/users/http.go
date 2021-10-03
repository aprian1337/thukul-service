package users

import (
	"aprian1337/thukul-service/business/users"
	"aprian1337/thukul-service/deliveries"
	"aprian1337/thukul-service/deliveries/users/requests"
	"aprian1337/thukul-service/deliveries/users/responses"
	"aprian1337/thukul-service/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	userUsecase users.Usecase
}

func NewUserController(uc users.Usecase) *Controller {
	return &Controller{
		userUsecase: uc,
	}
}

func (ctrl *Controller) GetUsersController(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := ctrl.userUsecase.GetAll(ctxNative)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromUsersListDomain(data))
}

func (ctrl *Controller) GetDetailUserController(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("id")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, errConvInt)
	}
	data, err := ctrl.userUsecase.GetById(ctxNative, convInt)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromUsersDomain(data))
}

func (ctrl *Controller) CreateUserController(c echo.Context) error {
	request := requests.UserRegister{}
	//u := middlewares.GetClaimUser(c)
	//fmt.Println(u)
	var err error
	err = c.Bind(&request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	ctxNative := c.Request().Context()
	var data users.Domain
	data, err = ctrl.userUsecase.Create(ctxNative, request.ToDomain())
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromDomainToCreateResponse(data))
}

func (ctrl *Controller) LoginUserController(c echo.Context) error {
	var login users.Domain
	var err error
	var token string
	ctxNative := c.Request().Context()

	request := requests.UserLogin{}
	err = c.Bind(&request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	login, token, err = ctrl.userUsecase.Login(ctxNative, request.Email, request.Password)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromUsersDomainToLogin(login, token))
}

func (cl *Controller) UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	convId, err := helpers.StringToUint(id)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := requests.UserRegister{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := cl.userUsecase.Update(ctx, req.ToDomain(), convId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromUsersDomain(data))
}

func (cl *Controller) DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	convId, err := helpers.StringToUint(id)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = cl.userUsecase.Delete(ctx, convId)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, convId)
}
