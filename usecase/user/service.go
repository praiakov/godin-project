package user

import (
	"errors"
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

//CreateUser Create an user
func (s *Service) GetUser(id entity.ID) (*entity.User, error) {
	return s.repo.Get(id)
}

//DeleteUser Delete an user
func (s *Service) DeleteUser(id entity.ID) error {
	u, err := s.GetUser(id)

	if u == nil {
		return errors.New("Not found")
	}

	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
