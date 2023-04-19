package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() {
	fmt.Println("DBへ接続を行います")
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_ca?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("DBへ接続しました")
}
