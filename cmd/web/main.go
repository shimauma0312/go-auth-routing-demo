package main

import (
	controller "go-auth-demo/pkg/controllers"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

func main() {
	app := &controller.App{
		Store: sessions.NewCookieStore([]byte("secret-key")),
	}

	router := app.NewRouter()
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
