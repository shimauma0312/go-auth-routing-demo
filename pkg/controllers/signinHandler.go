package controllers

import (
	tmp "go-auth-demo/constants"
	"net/http"
	"text/template"
)

func (app *App) signinHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(tmp.Signin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, tmp.Signin, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
