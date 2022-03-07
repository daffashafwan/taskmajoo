package routes

import (
	"github.com/daffashafwan/taskmajoo/app/middlewares"
	users "github.com/daffashafwan/taskmajoo/controllers/user"
	merchants "github.com/daffashafwan/taskmajoo/controllers/merchant"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig      middleware.JWTConfig
	UserController users.UserController
	MerchantController merchants.MerchantController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	//Users
	e.POST("users/login", cl.UserController.Login)

	//Merchants
	e.GET("users/:id/merchant", cl.MerchantController.GetByUserId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
}
