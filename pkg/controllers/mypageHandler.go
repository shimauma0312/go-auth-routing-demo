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
		UserId   string
	}{
		IsSignIn: app.IsUserLoggedIn(r),
		UserId:   sessions.Values["id"].(string),
	}
	if !data.IsSignIn {
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	t, err := template.ParseFiles(tmp.Layout, tmp.Mypage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
