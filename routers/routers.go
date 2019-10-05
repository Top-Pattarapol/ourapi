package routers

import (
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func GetUsers(c echo.Context) error {
	uu := []User{
		{
			ID:       1,
			Name:     "Leanne Graham",
			Username: "Bret",
			Email:    "Sincere@april.biz",
			Phone:    "1-770-736-8031 x56442",
		},
	}
	return c.JSON(http.StatusOK, uu)
}

func NewRouter() *echo.Echo {
	e := echo.New()
	e.GET("/users", GetUsers)
	return e
}