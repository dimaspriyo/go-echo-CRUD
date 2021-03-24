package main

import (
	"database/sql"
	"fmt"
	"go/internal/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	type Shared struct {
		ctx      echo.Context
		psqlconn *sql.DB
	}

	var shared Shared

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println(err.Error())
	}
	shared.psqlconn = db
	defer db.Close()

	e := echo.New()

	e.GET("/tet", func(c echo.Context) error {
		return c.String(http.StatusOK, "Your Real IP : "+c.RealIP())
	})

	e.GET("/aa", controller.SavePostgresql(shared))

	e.Start(":8000")
}
