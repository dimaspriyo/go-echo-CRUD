package postgresql

import (
	"go/internal/config"

	"github.com/labstack/echo/v4"
)

type PostgresqlController struct {
	s IPostgresqlService
}

func NewPostgresqlController(e *echo.Echo, shared config.GlobalShared, service IPostgresqlService) {
	controller := PostgresqlController{
		s: service,
	}

	g := e.Group("/postgresql", nil)
	g.GET("", controller.getAll)
	g.GET("/:id", controller.getById)
	g.POST("", controller.create)
	g.PUT("", controller.update)
	g.DELETE("", controller.delete)

}

func (c PostgresqlController) getAll(ctx echo.Context) error {
	return nil
}

func (c PostgresqlController) getById(ctx echo.Context) error {
	return nil
}

func (c PostgresqlController) create(ctx echo.Context) error {
	return nil
}

func (c PostgresqlController) update(ctx echo.Context) error {
	return nil
}

func (c PostgresqlController) delete(ctx echo.Context) error {
	return nil
}
