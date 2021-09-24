package users

import (
	"aprian1337/thukul-service/business/users"
	"aprian1337/thukul-service/deliveries"
	"aprian1337/thukul-service/deliveries/users/requests"
	"aprian1337/thukul-service/deliveries/users/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	userUsecase users.Usecase
}

func NewUserController(uc users.Usecase) *UserController {
	return &UserController{
		userUsecase: uc,
	}
}

func (ctrl *UserController) GetUsersController(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := ctrl.userUsecase.GetAll(ctxNative)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromUsersListDomain(data))
}

func (ctrl *UserController) GetDetailUserController(c echo.Context, id uint) error {
	ctxNative := c.Request().Context()
	data, err := ctrl.userUsecase.GetById(id, ctxNative)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromUsersDomain(data))
}

func (ctrl *UserController) CreateUserController(c echo.Context) error {
	request := requests.UserRegister{}
	err := c.Bind(&request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	ctxNative := c.Request().Context()
	data, err := ctrl.userUsecase.Create(ctxNative, request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, data)
}

//
//func (ctrl *UserController) LoginUserController(c echo.Context) error{
//
//}
