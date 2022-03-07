package routes

import (
	//"github.com/daffashafwan/taskmajoo/app/middlewares"
	users "github.com/daffashafwan/taskmajoo/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig      middleware.JWTConfig
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	//USERS
	e.POST("users/login", cl.UserController.Login)
}
