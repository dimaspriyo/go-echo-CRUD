package service

import (
	"go/internal/config"
	"go/internal/controller"
	"go/internal/repository"

	"github.com/labstack/echo/v4"
)

type IPostgresqlService interface {
	getbyId(id int, ctx echo.Context) controller.PostgresqlResponse
	getAll(ctx echo.Context) []controller.PostgresqlResponse
	create(req controller.PostgresqlRequest, ctx echo.Context) string
	update(id int, req controller.PostgresqlRequest, ctx echo.Context) string
	delete(id int, ctx echo.Context) string
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

func (s PostgresqlService) getbyId(id int, ctx echo.Context) controller.PostgresqlResponse {
	row := s.repo.FindById(id)
	resp := s.convertDAOtoDTO(row)
	return resp

}

func (s PostgresqlService) getAll(ctx echo.Context) []controller.PostgresqlResponse {

	var resp []controller.PostgresqlResponse
	rows := s.repo.FindAll()
	for _, v := range rows {
		temp := s.convertDAOtoDTO(v)
		resp = append(resp, temp)
	}

	return resp

}

func (s PostgresqlService) create(req controller.PostgresqlRequest, ctx echo.Context) string {
	ent := s.convertDTOtoDAO(req)
	rows, tx := s.repo.Create(ent, ctx)

	err := tx.Commit()
	if err != nil {
		panic(err.Error())
	}

	count, err := rows.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	if count != 1 {
		panic("Insert Failed")
	}

	return "Insert Success"

}

func (s PostgresqlService) update(id int, req controller.PostgresqlRequest, ctx echo.Context) string {

	ent := s.convertDTOtoDAO(req)
	row, tx := s.repo.Update(id, ent, ctx)

	err := tx.Commit()
	if err != nil {
		panic(err.Error())
	}

	count, err := row.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	if count != 1 {
		panic("Update Failed")
	}

	return "Update Success"

}

func (s PostgresqlService) delete(id int, ctx echo.Context) string {
	row, tx := s.repo.Delete(id, ctx)

	err := tx.Commit()
	if err != nil {
		panic(err.Error())
	}

	count, err := row.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	if count != 1 {
		panic("Delete Failed")
	}

	return "Delete Success"
}

func (s PostgresqlService) convertDTOtoDAO(req controller.PostgresqlRequest) repository.PostgresqlEntity {

	ent := repository.PostgresqlEntity{
		Name:    req.Name,
		Avatar:  req.Avatar,
		Address: req.Address,
	}

	return ent
}

func (s PostgresqlService) convertDAOtoDTO(ent repository.PostgresqlEntity) controller.PostgresqlResponse {
	resp := controller.PostgresqlResponse{
		Avatar:  ent.Avatar,
		Name:    ent.Name,
		Address: ent.Address,
	}

	return resp
}
