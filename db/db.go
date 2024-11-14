package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func NewMysqlStorage(cfg mysql.Config) *sql.DB {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CheckDBRun(db *sql.DB) error {
	return db.Ping()
}
