package routers

import (
	"github.com/labstack/echo"
	"github.com/ourapi/users"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.GET("/users", users.GetUsers)
	return e
}