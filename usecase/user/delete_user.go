package user

import (
	"errors"

	"github.com/praiakov/godin/entity"
)

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
