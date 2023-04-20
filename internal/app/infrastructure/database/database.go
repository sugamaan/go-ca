package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

func ConnectDb() {
	fmt.Println("DBへ接続を行います")

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err.Error())
	}

	// 環境変数から取得
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	addr := os.Getenv("MYSQL_ADDRESS")
	dbName := os.Getenv("MYSQL_DB_NAME")

	config := mysql.Config{
		User:      user,
		Passwd:    pass,
		Addr:      addr,
		DBName:    dbName,
		Loc:       jst,
		ParseTime: true,
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err.Error())
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("DBへ接続しました")
}
