package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"subscription-service/data"
	"testing"
)

var pages = []struct {
	name               string
	url                string
	expectedStatusCode int
	handler            http.HandlerFunc
	sessionData        map[string]any
	expectedHTML       string
}{
	{
		name:               "home",
		url:                "/",
		expectedStatusCode: http.StatusOK,
		handler:            testConfig.HomePage,
	},
	{
		name:               "login page",
		url:                "/login",
		expectedStatusCode: http.StatusSeeOther,
		handler:            testConfig.LoginPage,
		expectedHTML:       `<h1 class="mt-5">Login</h1>`,
	},
	{
		name:               "logout",
		url:                "/logout",
		expectedStatusCode: http.StatusOK,
		handler:            testConfig.LoginPage,
		expectedHTML:       `<h1 class="mt-5">Login</h1>`,
		sessionData: map[string]any{
			"userID": 1,
			"user":   data.User{},
		},
	},
}

func TestPages(t *testing.T) {
	pathToTemplates = "./templates"

	for _, p := range pages {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", p.url, nil)

		ctx := ctx(req)
		req = req.WithContext(ctx)

		if len(p.sessionData) > 0 {
			for k, v := range p.sessionData {
				testConfig.Session.Put(ctx, k, v)
			}
		}

		p.handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
		}

		if len(p.expectedHTML) > 0 {
			html := rr.Body.String()
			if !strings.Contains(html, p.expectedHTML) {
				t.Errorf("Returned wrong body: got %v want %v", html, p.expectedHTML)
			}
		}
	}
}

// go tool cover -html=coverage.out
func TestLogin(t *testing.T) {
	pathToTemplates = "./templates"

	data := url.Values{
		"email":    {"admin@example.com"},
		"password": {"abc123dfwpdcmsdi"},
	}

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(data.Encode()))
	ctx := ctx(req)
	req = req.WithContext(ctx)

	handler := http.HandlerFunc(testConfig.PostLoginPage)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("Returned wrong status code: got %v want %v", rr.Code, http.StatusSeeOther)
	}

	if !testConfig.Session.Exists(ctx, "userID") {
		t.Errorf("Session not created")
	}
}
