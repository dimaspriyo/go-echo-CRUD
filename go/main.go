package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"go/internal/repository"
)

func main() {

	repository.ESConnection()

	e := echo.New()

	e.GET("/tet", func(c echo.Context) error {
		return c.String(http.StatusOK, "Your Real IP : "+c.RealIP())
	})

	e.Start(":8000")
}
