package main

import (
	"fmt"
	"log"

	"go-ca/internal/app/infrastructure/database"

	"github.com/joho/godotenv"
)

func main() {
	// 環境変数の設定
	loadEnv()
	// DB接続
	database.ConnectDb()
}

func loadEnv() {
	if err := godotenv.Load("environments/local.env"); err != nil {
		log.Print("No .env file found")
	}
	fmt.Println("環境変数を読み込みました")
}
