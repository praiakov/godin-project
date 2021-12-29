package user

import (
	"github.com/praiakov/godin/entity"
)

//Writer user writer
type Writer interface {
	Create(e *entity.User) (entity.ID, error)
}

//Repository interface
type Repository interface {
	Writer
}

//UseCase interface
type UseCase interface {
	CreateUser(name, email string, totalMonth int) (entity.ID, error)
}
