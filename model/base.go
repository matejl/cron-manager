package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/spf13/viper"
	"strconv"
)

var dsn string

type base struct {
	Id int64 `json:"id"`
}

func GetConnection() *sql.DB {

	if dsn == "" {
		dsn = viper.GetString("db.username") + ":" +
			viper.GetString("db.password") + "@" +
			"tcp(" + viper.GetString("db.server") + ":" + strconv.Itoa(viper.GetInt("db.port")) + ")" +
			"/" + viper.GetString("db.database")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
