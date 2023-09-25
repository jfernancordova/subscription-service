package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.sessionLoad)

	mux.Get("/", app.home)

	mux.Get("/login", app.login)
	mux.Post("/login", app.postLogin)

	mux.Get("/logout", app.logout)

	mux.Get("/register", app.register)
	mux.Post("/register", app.postRegister)

	mux.Post("/activate-account", app.activateAccount)

	return mux
}
