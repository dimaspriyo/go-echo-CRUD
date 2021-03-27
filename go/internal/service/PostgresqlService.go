package service

import (
	"go/internal/config"
	"go/internal/controller"
	"go/internal/repository"

	"github.com/labstack/echo/v4"
)

type IPostgresqlService interface {
	getbyId(id int, ctx echo.Context)
	getAll(ctx echo.Context)
	create(req controller.PostgresqlRequest, ctx echo.Context)
	update(id int, req controller.PostgresqlRequest, ctx echo.Context)
	delete(id int, ctx echo.Context)
}

type PostgresqlService struct {
	shared config.GlobalShared
	repo   repository.IPostgresqlRepository
}

func NewPostgresqlService(s config.GlobalShared, r repository.PostgresqlRepository) IPostgresqlService {
	return PostgresqlService{
		shared: s,
		repo:   r,
	}
}

func (s PostgresqlService) getbyId(id int, ctx echo.Context) {

}

func (s PostgresqlService) getAll(ctx echo.Context) {

}

func (s PostgresqlService) create(req controller.PostgresqlRequest, ctx echo.Context) {

}

func (s PostgresqlService) update(id int, req controller.PostgresqlRequest, ctx echo.Context) {

}

func (s PostgresqlService) delete(id int, c echo.Context) {

}
