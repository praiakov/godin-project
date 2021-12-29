package user

import (
	"github.com/praiakov/godin/entity"
)

//CreateUser Create an user
func (s *Service) GetUser(id entity.ID) (*entity.User, error) {
	return s.repo.Get(id)
}
