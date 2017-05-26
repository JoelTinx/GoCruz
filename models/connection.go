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

func getConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		panic(err)
	}
	return db, err
}
