package presenter

import (
	"time"

	"github.com/praiakov/godin/entity"
)

//User data
type User struct {
	ID         entity.ID `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	TotalMonth int       `json:"total_month"`
}

type UserDetail struct {
	ID         entity.ID `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	PaidDate   time.Time `json:"paid_date"`
	DueDate    time.Time `json:"due_date"`
	TotalMonth int       `json:"total_month"`
}
