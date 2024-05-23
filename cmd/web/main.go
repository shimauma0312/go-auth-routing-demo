package main

import (
	controller "go-auth-demo/pkg/controllers"
	"go-auth-demo/pkg/models"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

func main() {
	// .envファイルから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数からセッションキーを取得する
	secret := os.Getenv("SESSION_KEY")
	if secret == "" {
		log.Fatal("SESSION_KEY is required")
	}

	// セッションキーを使用して新しいクッキーストアを作成する
	store := sessions.NewCookieStore([]byte(secret))
	// クッキーストアのオプションを設定する
	store.Options = &sessions.Options{
		Path:     "/",      // クッキーのパスを設定する
		MaxAge:   3600 * 8, // クッキーの有効期限を設定する（ここでは8時間）
		HttpOnly: true,     // JavaScriptからのアクセスを禁止する
		Secure:   true,     // HTTPS上でのみクッキーを送信する
	}

	// データベースに接続する
	db, err := models.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// 新しいアプリケーションを作成する
	app := &controller.App{
		Store: store,
		DB:    db,
	}

	// アプリケーションのルーターを作成する
	router := app.NewRouter()
	// HTTPサーバーを起動する
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
