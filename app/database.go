package app

import (
	"Golang-RestAPI/helper"
	"database/sql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:kaliberjunior1@tcp(localhost:3306)/category")
	helper.Panic(err)

	return db
}
