package usecase

import "github.com/IkezawaYuki/gopictweet/src/domain"

type UserRepository interface {
	FindAll() (*domain.Users, error)
	FindByEmail(string) (*domain.User, error)
	Create(*domain.User) (*domain.User, error)
	Update(*domain.User) (*domain.User, error)
	Delete(*domain.User) error
	DeleteAll() error
}
