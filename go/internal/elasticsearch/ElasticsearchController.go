package elasticsearch

import (
	"go/internal/config"

	"github.com/labstack/echo/v4"
)

type ElasticsearchController struct {
	Shared *config.GlobalShared
}

func NewElasticsearchController(e *echo.Echo, shared *config.GlobalShared) *echo.Echo{
	controller := ElasticsearchController{
		Shared: shared,
	}

	g := e.Group("/elasticsearch")
	g.
}


func (c *ElasticsearchController) getAll(ctx echo.Context) error{

}


func (c *ElasticsearchController) getById(ctx echo.Context) error{

}

func (c *ElasticsearchController) create(ctx echo.Context) error{

}

func (c *ElasticsearchController) update(ctx echo.Context) error{

}

func (c *ElasticsearchController) delete(ctx echo.Context) error{

}