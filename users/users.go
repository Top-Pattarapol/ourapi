package users

import (
"encoding/json"
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

type typicode struct {
}

func (tc *typicode) Decode(result interface{}) error {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return err
	}
	return json.NewDecoder(resp.Body).Decode(&result)
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
		&typicode{},
	}
	return api.getUsers(c)
}