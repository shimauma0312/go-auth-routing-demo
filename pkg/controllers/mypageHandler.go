package controllers

import (
	tmp "go-auth-demo/constants"
	"net/http"
	"text/template"
)

func (app *App) mypageHandler(w http.ResponseWriter, r *http.Request) {
	sessions, _ := app.Get(r, "user")

	data := struct {
		IsSignIn bool
		userId   string
	}{
		IsSignIn: sessions.Values["id"] != "",
		userId:   sessions.Values["id"].(string),
	}

	t, err := template.ParseFiles(tmp.Dir + tmp.Home)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, tmp.Home, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
