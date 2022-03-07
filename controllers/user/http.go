package users

import (
	"net/http"

	"github.com/daffashafwan/taskmajoo/business/users"
	"github.com/daffashafwan/taskmajoo/controllers/user/requests"
	"github.com/daffashafwan/taskmajoo/controllers/user/responses"
	"github.com/daffashafwan/taskmajoo/helpers/response"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase users.Usecase
}

func NewUserController(userUseCase users.Usecase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (userController UserController) Login(c echo.Context) error {

	userLogin := requests.UserLogin{}
	c.Bind(&userLogin)
	ctx := c.Request().Context()
	user, err := userController.UserUseCase.Login(ctx, userLogin.ToDomain())

	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(user))
}
