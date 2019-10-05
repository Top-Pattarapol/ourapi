package main

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func main()  {

	e := echo.New()
	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"hello": "hi"})
	})
	log.Fatal(e.Start(":1234"))
}