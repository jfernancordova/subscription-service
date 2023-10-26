package main

import (
	"fmt"
	"html/template"
	"net/http"
	"subscription-service/data"
	"time"
)

var pathToTemplates = "./cmd/web/templates"

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap     map[string]string
	IntMap        map[string]int
	FloatMap      map[string]float64
	Data          map[string]any
	Flash         string
	Warning       string
	Error         string
	Authenticated bool
	Now           time.Time
	User          *data.User
}

func (app *config) render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) {
	partials := []string{
		fmt.Sprintf("%s/base.layout.gohtml", pathToTemplates),
		fmt.Sprintf("%s/header.partial.gohtml", pathToTemplates),
		fmt.Sprintf("%s/navbar.partial.gohtml", pathToTemplates),
		fmt.Sprintf("%s/footer.partial.gohtml", pathToTemplates),
		fmt.Sprintf("%s/alerts.partial.gohtml", pathToTemplates),
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("%s/%s", pathToTemplates, t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	if td == nil {
		td = &TemplateData{}
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		app.errorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, app.AddDefaultData(td, r)); err != nil {
		app.errorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// AddDefaultData adds default data to the template data
func (app *config) AddDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Flash = app.session.PopString(r.Context(), "flash")
	td.Warning = app.session.PopString(r.Context(), "warning")
	td.Error = app.session.PopString(r.Context(), "error")
	if app.IsAuthenticated(r) {
		td.Authenticated = true
		user, ok := app.session.Get(r.Context(), "user").(data.User)
		if !ok {
			app.errorLog.Println("can't get user from session")
		} else {
			app.infoLog.Println("user recovered from session")
			td.User = &user
		}
	}
	td.Now = time.Now()

	return td
}

// IsAuthenticated checks if the user is logged in
func (app *config) IsAuthenticated(r *http.Request) bool {
	return app.session.Exists(r.Context(), "userID")
}
