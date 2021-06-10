package elasticsearch

import (
	"go/internal/config"

	"github.com/labstack/echo/v4"
)

type IElasticsearchService interface {
	getByName(name string, ctx echo.Context)

	save(name string, ctx echo.Context)
}

func newElasticsearchService(s *config.GlobalShared) IElasticsearchService {

	return ElasticsearchService{
		shared: s,
	}

}

type ElasticsearchService struct {
	shared *config.GlobalShared
	repo   IElasticsearchRepository
}

func (s ElasticsearchService) getByName(name string, ctx echo.Context) {

}

func (s ElasticsearchService) save(name string, ctx echo.Context) {

}
