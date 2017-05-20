package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbHost = "localhost"
	dbName = "dbGemace"
	dbUser = "root"
	dbPass =  "lca2012"//"123456"
	dbPort = 3306
)

var connsql string

func init() {
	connsql = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
}

func getConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", connsql)
	if err != nil {
		panic(err)
	}
	return db, err
}
