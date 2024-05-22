package controllers

import (
	tmp "go-auth-demo/constants"
	"net/http"
	"text/template"
)

func (app *App) signupHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(tmp.Dir + tmp.Signup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, tmp.Signup, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
