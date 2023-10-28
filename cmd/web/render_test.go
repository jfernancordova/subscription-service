package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestAddDefaultData tests the AddDefaultData function
func TestAddDefaultData(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := ctx(req)

	req = req.WithContext(ctx)

	testConfig.Session.Put(ctx, "flash", "flash")
	testConfig.Session.Put(ctx, "warning", "warning")
	testConfig.Session.Put(ctx, "error", "error")

	td := testConfig.AddDefaultData(&TemplateData{}, req)

	if td.Flash != "flash" {
		t.Errorf("flash value of %s is not expected value", td.Flash)
	}

	if td.Warning != "warning" {
		t.Errorf("warning value of %s is not expected value", td.Warning)
	}

	if td.Error != "error" {
		t.Errorf("error value of %s is not expected value", td.Error)
	}
}

func TestIsAuthenticated(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := ctx(req)
	req = req.WithContext(ctx)

	auth := testConfig.IsAuthenticated(req)
	if auth {
		t.Error("true is authenticated but IsAuthenticated returned false")
	}

	testConfig.Session.Put(ctx, "userID", 1)

	auth = testConfig.IsAuthenticated(req)
	if !auth {
		t.Errorf("false is not authenticated but IsAuthenticated returned true")
	}
}

func TestRender(t *testing.T) {
	pathToTemplates = "./templates"

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := ctx(req)
	req = req.WithContext(ctx)

	testConfig.render(rr, req, "home.page.gohtml", &TemplateData{})

	if rr.Code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rr.Code)
	}
}
