package postgresql

import (
	"database/sql"
	"go/internal/config"

	"github.com/labstack/echo/v4"
)

type IPostgresqlRepository interface {
	FindById(id int) PostgresqlEntity
	FindAll() []PostgresqlEntity
	Create(p PostgresqlEntity, ctx echo.Context) (sql.Result, *sql.Tx)
	Update(id int, p PostgresqlEntity, ctx echo.Context) (sql.Result, *sql.Tx)
	Delete(id int, ctx echo.Context) (sql.Result, *sql.Tx)
}

type PostgresqlRepository struct {
	shared config.GlobalShared
	entity PostgresqlEntity
	ctx    echo.Context
}

func NewPostgresqlRepository() IPostgresqlRepository {
	return PostgresqlRepository{}
}

func (r PostgresqlRepository) FindAll() []PostgresqlEntity {

	var res []PostgresqlEntity

	rows, err := r.shared.Psqlconn.Query("SELECT name, address, avatar FROM users")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		var temp PostgresqlEntity
		err := rows.Scan(&temp.Name, &temp.Address, &temp.Avatar)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, temp)
	}

	return res
}

func (r PostgresqlRepository) FindById(id int) PostgresqlEntity {
	var res PostgresqlEntity

	row := r.shared.Psqlconn.QueryRow(`SELECT name, address, avatar FROM users WHERE id=$1`, id)

	err := row.Scan(&res.Name, &res.Address, &res.Avatar)
	if err != nil {
		panic(err.Error())
	}

	return res
}

func (r PostgresqlRepository) Create(p PostgresqlEntity, ctx echo.Context) (sql.Result, *sql.Tx) {
	tx, err := r.shared.Psqlconn.BeginTx(r.ctx.Request().Context(), nil)
	if err != nil {
		panic(err.Error())
	}

	rows, err := tx.ExecContext(r.ctx.Request().Context(), `INSERT INTO users(name, address, avatar) VALUES($1,$2,$3)`, p.Name, p.Address, p.Avatar)
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}

	return rows, tx

}

func (r PostgresqlRepository) Update(id int, p PostgresqlEntity, ctx echo.Context) (sql.Result, *sql.Tx) {
	tx, err := r.shared.Psqlconn.BeginTx(r.ctx.Request().Context(), nil)
	if err != nil {
		panic(err.Error())
	}

	_ = r.FindById(id)

	rows, err := tx.ExecContext(r.ctx.Request().Context(), `UPDATE users SET name=$1, address=$2, avatar=$3 WHERE id=$4`, p.Name, p.Address, p.Avatar, id)
	if err != nil {
		panic(err.Error())
	}

	return rows, tx

}

func (r PostgresqlRepository) Delete(id int, ctx echo.Context) (sql.Result, *sql.Tx) {
	tx, err := r.shared.Psqlconn.BeginTx(r.ctx.Request().Context(), nil)
	if err != nil {
		panic(err.Error())
	}

	_ = r.FindById(id)

	rows, err := tx.ExecContext(r.ctx.Request().Context(), `DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		panic(err.Error())
	}

	return rows, tx
}
