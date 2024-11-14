package main

import (
	"github.com/go-sql-driver/mysql"
	"github.com/yann-fk-21/blogging/cmd/api"
	"github.com/yann-fk-21/blogging/db"
	"github.com/yann-fk-21/blogging/logger"
)

func main() {

	errLogger := logger.InitLogger()
	mysqlDB := db.NewMysqlStorage(mysql.Config{
		// TODO
	})

	err := db.CheckDBRun(mysqlDB)
	if err != nil {
		errLogger.Println(err)
		errLogger.Fatal(err)
	}

	server := api.NewServer(":8080", mysqlDB)
	err = server.Run()
	if err != nil {
		errLogger.Println(err)
		errLogger.Fatal(err)
	}
}
