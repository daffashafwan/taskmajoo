package routes

import (
	"github.com/daffashafwan/taskmajoo/app/middlewares"
	users "github.com/daffashafwan/taskmajoo/controllers/user"
	merchants "github.com/daffashafwan/taskmajoo/controllers/merchant"
	transactions "github.com/daffashafwan/taskmajoo/controllers/transaction"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig      middleware.JWTConfig
	UserController users.UserController
	MerchantController merchants.MerchantController
	TransactionController transactions.TransactionController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	//Users
	e.POST("users/login", cl.UserController.Login)

	//Merchants
	e.GET("users/:id/merchant", cl.MerchantController.GetByUserId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)

	//Transaction
	e.GET("users/:id/transaction", cl.TransactionController.GetByMerchantId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("users/:id/transaction/outlet", cl.TransactionController.GetByMerchantIdWithOutlet, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
}
