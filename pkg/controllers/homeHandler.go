package controllers

import (
	tmp "go-auth-demo/constants"
	"net/http"
	"text/template"
)

func (app *App) homeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		IsSignIn bool
	}{
		IsSignIn: app.IsUserLoggedIn(r),
	}

	t, err := template.ParseFiles(tmp.Layout, tmp.Home)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
