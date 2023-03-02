package app

import (
	"database/sql"
	"fmt"

	"github.com/Alfeenn/article/helper"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golangdb")
	helper.PanicIfErr(err)
	fmt.Println("Connected")
	return db
}
