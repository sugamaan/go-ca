package main

import (
	"fmt"
	"go-ca/internal/app/infrastructure/register"
	"log"
	"net/http"

	"go-ca/internal/app/infrastructure/database"

	"github.com/joho/godotenv"
)

func main() {
	// 環境変数の設定
	loadEnv()

	// DB接続
	database.ConnectDb()

	// サーバー
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/tasks", register.Tasks)

	// サーバーの起動
	fmt.Println("Starting server on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func loadEnv() {
	if err := godotenv.Load("environments/local.env"); err != nil {
		log.Print("No .env file found")
	}
	fmt.Println("環境変数を読み込みました")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "正常に疎通")
}
