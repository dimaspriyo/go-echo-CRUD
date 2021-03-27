package repository

import (
	"go/internal/config"

	"github.com/labstack/echo/v4"
)

type IPostgresqlRepository interface {
	getbyId(id int) PostgresqlEntity
	getAll() []PostgresqlEntity
	create(p PostgresqlEntity)
	update(id int, p PostgresqlEntity)
	delete(id int)
}

type PostgresqlRepository struct {
	shared config.GlobalShared
	entity PostgresqlEntity
	ctx    echo.Context
}

func NewPostgresqlRepository(s config.GlobalShared, e PostgresqlEntity, c echo.Context) IPostgresqlRepository {
	return PostgresqlRepository{
		shared: s,
		entity: e,
		ctx:    c,
	}
}

func (r PostgresqlRepository) getAll() []PostgresqlEntity {

	var res []PostgresqlEntity

	rows, err := r.shared.Psqlconn.Query("SELECT name, address, avatar FROM users")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		var temp PostgresqlEntity
		err := rows.Scan(&temp.name, &temp.address, &temp.avatar)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, temp)
	}

	return res
}

func (r PostgresqlRepository) getbyId(id int) PostgresqlEntity {
	var res PostgresqlEntity

	row := r.shared.Psqlconn.QueryRow(`SELECT name, address, avatar FROM users WHERE id=$1`, id)

	err := row.Scan(&res.name, &res.address, &res.avatar)
	if err != nil {
		panic(err.Error())
	}

	return res
}

func (r PostgresqlRepository) create(p PostgresqlEntity) {
	tx, err := r.shared.Psqlconn.BeginTx(r.ctx.Request().Context(), nil)
	if err != nil {
		panic(err.Error())
	}

	rows, err := tx.ExecContext(r.ctx.Request().Context(), `INSERT INTO users(name, address, avatar) VALUES($1,$2,$3)`, p.name, p.address, p.avatar)
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}

	err = tx.Commit()
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

}

func (r PostgresqlRepository) update(id int, p PostgresqlEntity) {
	tx, err := r.shared.Psqlconn.BeginTx(r.ctx.Request().Context(), nil)
	if err != nil {
		panic(err.Error())
	}

	_ = r.getbyId(id)

	rows, err := tx.ExecContext(r.ctx.Request().Context(), `UPDATE users SET name=$1, address=$2, avatar=$3 WHERE id=$4`, p.name, p.address, p.avatar, id)
	if err != nil {
		panic(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		panic(err.Error())
	}

	count, err := rows.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	if count != 1 {
		panic("Update Failed")
	}

}

func (r PostgresqlRepository) delete(id int) {
	tx, err := r.shared.Psqlconn.BeginTx(r.ctx.Request().Context(), nil)
	if err != nil {
		panic(err.Error())
	}

	_ = r.getbyId(id)

	rows, err := tx.ExecContext(r.ctx.Request().Context(), `DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		panic(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		panic(err.Error())
	}

	count, err := rows.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	if count != 1 {
		panic("Delete Failed")
	}
}
