package main

import "net/http"

func (app *config) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.gohtml", nil)
}

func (app *config) login(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.gohtml", nil)
}

func (app *config) postLogin(w http.ResponseWriter, r *http.Request) {
}

func (app *config) logout(w http.ResponseWriter, r *http.Request) {

}

func (app *config) register(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "register.page.gohtml", nil)
}

func (app *config) postRegister(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "register.page.gohtml", nil)
}

func (app *config) activateAccount(w http.ResponseWriter, r *http.Request) {
}
