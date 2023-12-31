package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

var routes = []string{
	"/",
	"/login",
	"/logout",
	"/register",
	"/activate",
	"/members/plans",
	"/members/subscribe",
}

// TestRoutesExist tests that all routes exist
func TestRoutesExist(t *testing.T) {
	testRoutes := testConfig.routes()
	chiRoutes := testRoutes.(chi.Router)

	for _, route := range routes {
		found := false
		_ = chi.Walk(chiRoutes, func(method string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
			if route == foundRoute {
				found = true
			}
			return nil
		})

		if !found {
			t.Errorf("route %s not found", route)
		}
	}
}
