package postgresql

import (
	"go/internal/config"
	"go/internal/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostgresqlController struct {
	s IPostgresqlService
}

func NewPostgresqlController(e *echo.Echo, shared *config.GlobalShared) {
	controller := PostgresqlController{
		s: NewPostgresqlService(shared),
	}

	g := e.Group("/postgresql")
	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	g.GET("", controller.getAll)
	g.GET("/:id", controller.getById)
	g.POST("", controller.create)
	g.PUT("", controller.update)
	g.DELETE("", controller.delete)

}

func (c *PostgresqlController) getAll(ctx echo.Context) error {
	var response MainResponse

	res := c.s.GetAll(ctx)

	response.Data = res
	response.Success = true

	return ctx.JSON(http.StatusOK, response)
}

func (c PostgresqlController) getById(ctx echo.Context) error {
	var response MainResponse

	id := util.GetIDInt64Param(ctx)
	res := c.s.GetbyId(id, ctx)

	response.Data = res
	response.Success = true

	return ctx.JSON(http.StatusOK, response)
}

func (c PostgresqlController) create(ctx echo.Context) error {
	var response MainResponse
	var req PostgresqlRequest

	err := ctx.Bind(&req)
	if err != nil {
		panic(err.Error())
	}

	res := c.s.Create(req, ctx)

	response.Data = res
	response.Success = true

	return ctx.JSON(http.StatusCreated, response)
}

func (c PostgresqlController) update(ctx echo.Context) error {
	var response MainResponse
	var req PostgresqlRequest

	id := util.GetIDInt64Param(ctx)
	err := ctx.Bind(ctx.Request())
	if err != nil {
		panic(err.Error())
	}
	res := c.s.Update(id, req, ctx)

	response.Data = res
	response.Success = true

	return ctx.JSON(http.StatusOK, response)
}

func (c PostgresqlController) delete(ctx echo.Context) error {
	var response MainResponse

	id := util.GetIDInt64Param(ctx)
	res := c.s.Delete(id, ctx)

	response.Data = res
	response.Success = true

	return ctx.JSON(http.StatusOK, response)
}
