package main

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func main()  {
	users := []Users{
		{1,"Leanne Graham","Bret", "Sincere@april.biz", "1-770-736-8031 x56442"},
	}

	e := echo.New()
	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, users)
	})
	log.Fatal(e.Start(":1234"))
}