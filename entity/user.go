package entity

import (
	"errors"
	"time"
)

//User data
type User struct {
	ID         ID
	Name       string
	Email      string
	PaidDate   time.Time
	DueDate    time.Time
	TotalMonth int
}

//Newuser
//Create a new User
func NewUser(name, email string, paidDate, dueDate time.Time, totalMonth int) *User {
	return &User{
		ID:         NewID(),
		Name:       name,
		Email:      email,
		PaidDate:   paidDate,
		DueDate:    dueDate,
		TotalMonth: totalMonth,
	}
}

//Validate validate data
func (u *User) Validate() error {
	if u.Name == "" || u.Email == "" || u.TotalMonth == 0 {
		return errors.New("validate")
	}

	return nil
}
