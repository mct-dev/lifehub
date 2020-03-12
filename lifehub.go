package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", home)

	e.Logger.Fatal(e.Start(":1323"))
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hey there breh")
}
