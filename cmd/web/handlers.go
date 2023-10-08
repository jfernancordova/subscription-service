package main

import "net/http"

// HomePage displays the home page
func (app *config) HomePage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.gohtml", nil)
}

// LoginPage displays the login page
func (app *config) LoginPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.gohtml", nil)
}

// PostLoginPage handles the login form submission
func (app *config) PostLoginPage(w http.ResponseWriter, r *http.Request) {
	_ = app.session.RenewToken(r.Context())

	// parse form post
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Println(err)
	}

	// get email and password from form post
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := app.models.User.GetByEmail(email)
	if err != nil {
		app.session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// check password
	validPassword, err := user.PasswordMatches(password)
	if err != nil {
		app.session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if !validPassword {
		app.session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// add user to session
	app.session.Put(r.Context(), "userID", user.ID)
	app.session.Put(r.Context(), "user", user)

	app.session.Put(r.Context(), "flash", "Successful login!")

	// redirect the user
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Logout logs the user out
func (app *config) Logout(w http.ResponseWriter, r *http.Request) {
	// clean up session
	_ = app.session.Destroy(r.Context())
	_ = app.session.RenewToken(r.Context())

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// RegisterPage displays the register page
func (app *config) RegisterPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "register.page.gohtml", nil)
}

// PostRegisterPage handles the register form submission
func (app *config) PostRegisterPage(w http.ResponseWriter, r *http.Request) {
	// create a user

	// send an activation email

	// subscbribe the user to an account
}

// ActivateAccount activates a user's account
func (app *config) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	// validate url

	// generate an invoice

	// send an email with attachments

	// send an email with the invoice attached
}
