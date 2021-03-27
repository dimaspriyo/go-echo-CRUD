package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type GlobalShared struct {
	Psqlconn *sql.DB
}

type PsqlConn struct {
	Host     string
	Port     int64
	User     string
	Password string
	Dbname   string
}

var Shared GlobalShared

func InitShared(psql PsqlConn) GlobalShared {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", psql.Host, psql.Port, psql.User, psql.Password, psql.Dbname)
	psqldb, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer psqldb.Close()

	return GlobalShared{
		Psqlconn: psqldb,
	}
}
