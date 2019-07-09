package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", HelloWorld())
	e.Logger.Fatal(e.Start(":8000"))
}

type H map[string]interface{}

func HelloWorld() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, H{
			"code": http.StatusOK,
			"data": "Hello World",
		})
	}
}
