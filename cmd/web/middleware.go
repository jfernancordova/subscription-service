package main

import "net/http"

func (app *config) sessionLoad(next http.Handler) http.Handler {
	return app.session.LoadAndSave(next)
}
