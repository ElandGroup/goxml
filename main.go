package main

import (
	"goxml/logs"
	"net/http"
	//"gorouter/api"
	"github.com/labstack/echo"
)

var (
	e = echo.New()
)

func main() {

	e.GET("/", func(c echo.Context) error {
		logs.Warning.Println("hello good")
		return c.String(http.StatusOK, "Hello echo api")
	})
	RegisterUserApi()
	logs.Warning.Println("RegisterUserApi end")

	e.Start(":1111")
}
