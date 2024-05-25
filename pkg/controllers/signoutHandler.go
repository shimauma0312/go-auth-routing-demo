package controllers

import "net/http"

func (app *App) signoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := app.Get(r, "user")
	session.Values["id"] = nil
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
