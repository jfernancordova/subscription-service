package main

import "net/http"

// SessionLoad adds CSRF protection to all POST requests
func (app *config) SessionLoad(next http.Handler) http.Handler {
	return app.session.LoadAndSave(next)
}
