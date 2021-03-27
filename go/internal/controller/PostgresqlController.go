package controller

import (
	"go/internal/config"
	"go/internal/service"

	"github.com/labstack/echo/v4"
)

type PostgresqlController struct {
	s service.IPostgresqlService
}

func NewPostgresqlController(e *echo.Echo, shared config.GlobalShared, service service.IPostgresqlService) {
	controller := PostgresqlController{
		s: service,
	}

	g := e.Group("/postgresql", nil)
	g.GET("", controller.getAll())

}

func (c PostgresqlController) getAll(ctx echo.Context) {
}

func (c PostgresqlController) geyById(ctx echo.Context) {
}

func (c PostgresqlController) create(ctx echo.Context) {
}

func (c PostgresqlController) update(ctx echo.Context) {
}

func (c PostgresqlController) delete(ctx echo.Context) {
}
