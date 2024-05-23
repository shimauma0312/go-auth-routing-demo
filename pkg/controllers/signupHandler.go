package controllers

import (
	tmp "go-auth-demo/constants"
	"log"
	"net/http"
	"text/template"
)

func (app *App) signupHandler(w http.ResponseWriter, r *http.Request) {
	if app.DB == nil {
		log.Fatal("Database instance is not initialized")
	}
	if r.Method == "GET" {
		t, err := template.ParseFiles(tmp.Dir + tmp.Signup)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = t.ExecuteTemplate(w, tmp.Signup, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		formUsername := r.FormValue("username")
		formPassword := r.FormValue("password")
		err := app.DB.CreateUser(formUsername, formPassword)
		if err != nil {
			http.Error(w, "An error occurred", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/signin", http.StatusFound)
	}
}
