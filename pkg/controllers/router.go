package controllers

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type App struct {
	Store *sessions.CookieStore
}

func (app *App) Get(r *http.Request, name string) (*sessions.Session, error) {
	return app.Store.Get(r, name)
}

func (app *App) NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/home", app.homeHandler)
	mux.HandleFunc("/signin", app.signinHandler)
	mux.HandleFunc("/signgout", app.signoutHandler)
	mux.HandleFunc("/signup", app.signupHandler)
	mux.HandleFunc("/Mypage", app.mypageHandler)

	return mux
}
