package data

import (
	"database/sql"
	"fmt"
	"time"
)

// TestNew is the function used to create an instance of the data package. It returns the type
func TestNew(dbPool *sql.DB) Models {
	db = dbPool
	return Models{
		User: &UserMocked{},
		Plan: &PlanMocked{},
	}
}

// UserMocked is the type for users
type UserMocked struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	Password  string
	Active    int
	IsAdmin   int
	CreatedAt time.Time
	UpdatedAt time.Time
	Plan      *Plan
}

// GetAll returns all users
func (u *UserMocked) GetAll() ([]*User, error) {
	var users []*User
	user := User{
		ID:        1,
		Email:     "admin@example.com",
		FirstName: "Admin",
		LastName:  "Admin",
		Password:  "verysecret",
		Active:    1,
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	users = append(users, &user)
	return users, nil
}

// GetByEmail returns one user by email
func (u *UserMocked) GetByEmail(email string) (*User, error) {
	user := User{
		ID:        1,
		Email:     "admin@example.com",
		FirstName: "Admin",
		LastName:  "Admin",
		Password:  "verysecret",
		Active:    1,
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &user, nil
}

// GetOne returns one user by id
func (u *UserMocked) GetOne(id int) (*User, error) {
	return u.GetByEmail("")
}

// Update updates a user
func (u *UserMocked) Update(user User) error {
	return nil
}

// Delete deletes a user
func (u *UserMocked) Delete() error {
	return nil
}

// DeleteByID deletes a user by id
func (u *UserMocked) DeleteByID(id int) error {
	return nil
}

// Insert inserts a user
func (u *UserMocked) Insert(user User) (int, error) {
	return 2, nil
}

// ResetPassword resets a user's password
func (u *UserMocked) ResetPassword(password string) error {
	return nil
}

// PasswordMatches checks if a user's password matches
func (u *UserMocked) PasswordMatches(plainText string) (bool, error) {
	return true, nil
}

// PlanMocked is the type for plans
type PlanMocked struct {
	ID                  int
	PlanName            string
	PlanAmount          int
	PlanAmountFormatted string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

// GetAll returns all plans
func (p *PlanMocked) GetAll() ([]*Plan, error) {
	var plans []*Plan
	plan := Plan{
		ID:         1,
		PlanName:   "Test Plan",
		PlanAmount: 1000,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	plans = append(plans, &plan)
	return plans, nil
}

// GetOne returns one plan by id
func (p *PlanMocked) GetOne(id int) (*Plan, error) {
	plan := Plan{
		ID:         1,
		PlanName:   "Test Plan",
		PlanAmount: 1000,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	return &plan, nil
}

// SubscribeUserToPlan subscribes a user to a plan
func (p *PlanMocked) SubscribeUserToPlan(user User, plan Plan) error {
	return nil
}

// AmountForDisplay returns the amount for display
func (p *PlanMocked) AmountForDisplay() string {
	amount := float64(p.PlanAmount) / 100.0
	return fmt.Sprintf("$%.2f", amount)
}
