package repository

import (
	"github.com/IkezawaYuki/pictweet-api/src/domain/model"
)

type UserRepository interface {
	FindAll() (*model.Users, error)
	FindByEmail(string) (*model.User, error)
	Create(*model.User) (*model.User, error)
	Update(*model.User) (*model.User, error)
	Delete(*model.User) error
	DeleteAll() error
	FindBySessionID(string) (*model.User, error)
	FindByUUID(string) (*model.User, error)
}
