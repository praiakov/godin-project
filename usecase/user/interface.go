package user

import (
	"github.com/praiakov/godin/entity"
)

//Writer user writer
type Writer interface {
	Create(e *entity.User) (entity.ID, error)
}

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.User, error)
}

//Repository interface
type Repository interface {
	Writer
	Reader
}

//UseCase interface
type UseCase interface {
	CreateUser(name, email string, totalMonth int) (entity.ID, error)
	GetUser(id entity.ID) (*entity.User, error)
}
