package users

import (
	"github.com/ourapi/typicode"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type Decoder interface {
	Decode(result interface{}) error
}

type usersApi struct {
	sevice Decoder
}

func (u *usersApi) getUsers(c echo.Context) error {
	uu := []User{}
	err := u.sevice.Decode(&uu)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, uu)
}

func GetUsers(c echo.Context) error {
	api := &usersApi{
		typicode.Get("/users"),
	}
	return api.getUsers(c)
}