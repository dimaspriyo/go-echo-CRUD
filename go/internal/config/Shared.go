package config

import (
	"database/sql"
	"fmt"
)

type GlobalShared struct {
	Psqlconn *sql.DB
}

var Shared GlobalShared

func InitShared() GlobalShared {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	psqldb, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer psqldb.Close()

	return GlobalShared{
		Psqlconn: psqldb,
	}
}
