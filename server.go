package main

import (
	"covid-tracker/src/api"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, mayank!")
	})
	e.GET("/cases",func(c echo.Context) error {
		return api.CasesHandler(c)
	})
	e.Logger.Fatal(e.Start(":1323"))

}
