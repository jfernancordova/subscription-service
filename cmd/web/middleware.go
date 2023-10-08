package main

import "net/http"

// SessionLoad adds CSRF protection to all POST requests
func (app *Config) SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}