package app

import (
	"Golang-RestAPI/helper"
	"database/sql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang-restapi")
	helper.Panic(err)

	return db
}
