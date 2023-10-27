package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"subscription-service/data"
	"time"

	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
)

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
		msg := Message{
			To:      email,
			Subject: "Failed log in attempt",
			Data:    "Invalid login attempt!",
		}

		app.sendEmail(msg)

		app.session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// add user to session
	app.session.Put(r.Context(), "userID", user.ID)
	app.session.Put(r.Context(), "user", user)

	app.session.Put(r.Context(), "flash", "Successful login!")

	// redirect the user
	http.Redirect(w, r, "/members/plans", http.StatusSeeOther)
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
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Println(err)
	}

	// TODO - validate data

	// create a user
	u := data.User{
		Email:     r.Form.Get("email"),
		FirstName: r.Form.Get("first-name"),
		LastName:  r.Form.Get("last-name"),
		Password:  r.Form.Get("password"),
		Active:    0,
		IsAdmin:   0,
	}

	_, err = u.Insert(u)
	if err != nil {
		app.session.Put(r.Context(), "error", "Failed to create user.")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	// send an activation email
	url := fmt.Sprintf("http://localhost/activate?email=%s", u.Email)
	signedURL := GenerateTokenFromString(url)
	app.infoLog.Println(signedURL)

	msg := Message{
		To:       u.Email,
		Subject:  "Activate your account",
		Template: "confirmation-email",
		Data:     template.HTML(signedURL),
	}
	app.sendEmail(msg)

	app.session.Put(r.Context(), "flash", "Your account has been created! Please check your email to activate your account.")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// ActivateAccount activates a user's account
func (app *config) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	// validate url
	url := r.RequestURI
	testURL := fmt.Sprintf("http://localhost%s", url)
	okay := VerifyToken(testURL)

	if !okay {
		app.session.Put(r.Context(), "error", "Invalid activation link.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// activate account
	u, err := app.models.User.GetByEmail(r.URL.Query().Get("email"))
	if err != nil {
		app.session.Put(r.Context(), "error", "No user found.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u.Active = 1
	err = u.Update()
	if err != nil {
		app.session.Put(r.Context(), "error", "Unable to update user.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	app.session.Put(r.Context(), "flash", "Your account has been activated! Please log in.")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *config) SubscribeToPlan(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	planID, err := strconv.Atoi(id)
	if err != nil {
		app.errorLog.Println(err)
	}

	plan, err := app.models.Plan.GetOne(planID)
	if err != nil {
		app.session.Put(r.Context(), "error", "Unable to find plan.")
		http.Redirect(w, r, "/members/plans", http.StatusSeeOther)
		return
	}

	user, ok := app.session.Get(r.Context(), "user").(data.User)
	if !ok {
		app.session.Put(r.Context(), "error", "Log in first.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	app.wait.Add(1)
	go func() {
		defer app.wait.Done()
		invoice, err := app.Invoice(user, plan)
		if err != nil {
			app.errorChan <- err
		}

		msg := Message{
			To:       user.Email,
			Subject:  "Your Invoice",
			Data:     invoice,
			Template: "invoice",
		}
		app.sendEmail(msg)
	}()

	app.wait.Add(1)
	go func() {
		defer app.wait.Done()
		pdf := app.Manual(user, plan)
		err := pdf.OutputFileAndClose(fmt.Sprintf("./tmp/%d_manual.pdf", user.ID))
		if err != nil {
			app.errorChan <- err
			return
		}

		msg := Message{
			To:      user.Email,
			Subject: "Your Manual",
			Data:    "Please find your manual attached.",
			AttachmentMap: map[string]string{
				"Manual.pdf": fmt.Sprintf("./tmp/%d_manual.pdf", user.ID),
			},
		}

		app.sendEmail(msg)
	}()

	// subscribe user to plan
	err = app.models.Plan.SubscribeUserToPlan(user, *plan)
	if err != nil {
		app.session.Put(r.Context(), "error", "Unable to subscribe to plan.")
		http.Redirect(w, r, "/members/plans", http.StatusSeeOther)
		return
	}

	u, err := app.models.User.GetOne(user.ID)
	if err != nil {
		app.session.Put(r.Context(), "error", "Unable to find user.")
		http.Redirect(w, r, "/members/plans", http.StatusSeeOther)
		return
	}

	app.session.Put(r.Context(), "user", u)

	// redirect the user
	app.session.Put(r.Context(), "flash", "Your subscription has been created!")
	http.Redirect(w, r, "/members/plans", http.StatusSeeOther)
}

func (app *config) Manual(u data.User, plan *data.Plan) *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "Letter", "")
	pdf.SetMargins(10, 13, 10)

	importer := gofpdi.NewImporter()

	time.Sleep(5 * time.Second)

	t := importer.ImportPage(pdf, "./pdf/manual.pdf", 1, "/MediaBox")
	pdf.AddPage()

	importer.UseImportedTemplate(pdf, t, 0, 0, 215.9, 0)

	pdf.SetX(75)
	pdf.SetY(150)

	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, 4, fmt.Sprintf("%s %s", u.FirstName, u.LastName), "", "C", false)
	pdf.Ln(5)
	pdf.MultiCell(0, 4, fmt.Sprintf("%s User Guide", plan.PlanName), "", "C", false)

	return pdf
}

func (app *config) Invoice(u data.User, plan *data.Plan) (string, error) {
	return plan.PlanAmountFormatted, nil
}

func (app *config) ChooseSubscription(w http.ResponseWriter, r *http.Request) {
	plans, err := app.models.Plan.GetAll()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	dataMap := make(map[string]any)
	dataMap["plans"] = plans

	app.render(w, r, "plans.page.gohtml", &TemplateData{
		Data: dataMap,
	})
}
