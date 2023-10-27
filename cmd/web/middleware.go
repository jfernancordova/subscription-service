package main

import "net/http"

// SessionLoad adds CSRF protection to all POST requests
func (app *config) SessionLoad(next http.Handler) http.Handler {
	return app.session.LoadAndSave(next)
}

// Auth checks if the user is logged in and redirects to the login page if they are not
func (app *config) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if the user is not authenticated, redirect them to the login page
		if !app.session.Exists(r.Context(), "userID") {
			app.session.Put(r.Context(), "error", "You must be logged in to do that")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// otherwise call the next handler
		next.ServeHTTP(w, r)
	})
}