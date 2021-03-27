package controller

import (
	"go/internal/service"

	"github.com/labstack/echo/v4"
)

type PostgresqlController struct {
	s service.IPostgresqlService
}

func NewPostgresqlController() {

}

func (c PostgresqlController) getAll(ctx echo.Context) {
}
