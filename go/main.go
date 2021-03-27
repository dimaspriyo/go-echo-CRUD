package main

import (
	"fmt"
	"go/internal/config"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var Shared config.GlobalShared

func init() {
	viper.SetConfigName("env") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./env") // config file path

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	psqlConn := config.PsqlConn{
		Host:     viper.GetString("postgresql.host"),
		Port:     viper.GetInt64("postgresql.port"),
		User:     viper.GetString("postgresql.user"),
		Password: viper.GetString("postgresql.password"),
		Dbname:   viper.GetString("postgresql.dbname"),
	}
	Shared = config.InitShared(psqlConn)

}

func main() {

	e := echo.New()

	e.GET("/tet", func(c echo.Context) error {
		return c.String(http.StatusOK, "Your Real IP : "+c.RealIP())
	})

	e.Start(":8000")
}
