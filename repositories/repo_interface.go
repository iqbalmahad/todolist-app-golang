package repositories

import (
	"github.com/iqbalmahad/todolist-app-golang.git/models"
)

type UserRepo interface {
	Create(user models.User) error
	Read(id string) (models.User, error)
	Update(user models.User) error
	Delete(id string) error
	Index() ([]models.User, error)
}

type TodoRepo interface {
	Create(todo models.Todo) error
	Read(id string) (models.Todo, error)
	Update(todo models.Todo) error
	Delete(id string) error
	Index() ([]models.Todo, error)
}

type AuthRepo interface {
	GetByUsernameOrEmail(usernameOrEmail string) (*models.User, error)
}
