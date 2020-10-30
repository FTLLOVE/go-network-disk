package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB
var err error

func init() {
	DB, err = sql.Open("mysql", "root:Abcdef@123456@tcp(localhost:3306)/netdisk?charset=utf8")
	if err != nil {
		log.Printf("sql.Open fail,err: %v \n", err.Error())
		return
	}
	if err = DB.Ping(); err != nil {
		log.Printf("connec fail,err: %b \n", err.Error())
	}
}
