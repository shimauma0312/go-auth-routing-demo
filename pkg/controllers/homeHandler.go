package controllers

import (
	tmp "go-auth-demo/constants"
	"net/http"
	"text/template"
)

func (app *App) homeHandler(w http.ResponseWriter, r *http.Request) {
	sessions, _ := app.Get(r, "user")

	data := struct {
		IsSignIn bool
	}{
		IsSignIn: sessions.Values["id"] != "",
	}

	t, err := template.ParseFiles(tmp.Dir + tmp.Home)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, tmp.Home, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
