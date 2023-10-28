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
		User: &UserTest{},
		Plan: &PlanTest{},
	}
}

// UserTest is the type for users
type UserTest struct {
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
func (u *UserTest) GetAll() ([]*User, error) {
	var users []*User
	user := User{
		ID:        1,
		Email:     "Admin",
		FirstName: "Admin",
		LastName:  "Admin",
		Password:  "Admin",
		Active:    1,
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	users = append(users, &user)
	return users, nil
}

// GetByEmail returns one user by email
func (u *UserTest) GetByEmail(email string) (*User, error) {
	user := User{
		ID:        1,
		Email:     "Admin",
		FirstName: "Admin",
		LastName:  "Admin",
		Password:  "Admin",
		Active:    1,
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &user, nil
}

// GetOne returns one user by id
func (u *UserTest) GetOne(id int) (*User, error) {
	return u.GetByEmail("")
}

// Update updates a user
func (u *UserTest) Update() error {
	return nil
}

// Delete deletes a user
func (u *UserTest) Delete() error {
	return nil
}

// DeleteByID deletes a user by id
func (u *UserTest) DeleteByID(id int) error {
	return nil
}

// Insert inserts a user
func (u *UserTest) Insert(user User) (int, error) {
	return 2, nil
}

// ResetPassword resets a user's password
func (u *UserTest) ResetPassword(password string) error {
	return nil
}

// PasswordMatches checks if a user's password matches
func (u *UserTest) PasswordMatches(plainText string) (bool, error) {
	return true, nil
}

// PlanTest is the type for plans
type PlanTest struct {
	ID                  int
	PlanName            string
	PlanAmount          int
	PlanAmountFormatted string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

// GetAll returns all plans
func (p *PlanTest) GetAll() ([]*Plan, error) {
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
func (p *PlanTest) GetOne(id int) (*Plan, error) {
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
func (p *PlanTest) SubscribeUserToPlan(user User, plan Plan) error {
	return nil
}

// AmountForDisplay returns the amount for display
func (p *PlanTest) AmountForDisplay() string {
	amount := float64(p.PlanAmount) / 100.0
	return fmt.Sprintf("$%.2f", amount)
}
