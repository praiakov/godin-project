package presenter

import (
	"github.com/praiakov/godin/entity"
)

//User data
type User struct {
	ID         entity.ID `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	TotalMonth int       `json:"total_month"`
}
