package user

import (
	"time"

	"github.com/praiakov/godin/entity"
)

//Service  interface
type Service struct {
	repo Repository
}

//NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateUser Create an user
func (s *Service) CreateUser(name, email string, totalMonth int) (entity.ID, error) {
	data := time.Now()

	e := entity.NewUser(name, email, data, data.AddDate(0, totalMonth, 0), totalMonth)

	return s.repo.Create(e)
}
