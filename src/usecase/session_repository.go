package usecase

import "github.com/IkezawaYuki/gopictweet/src/domain"

type SessionRepository interface {
	Create(*domain.Session) (*domain.Session, error)
	FindByUserID(string) (*domain.Session, error)
	Check(string) (bool, error)
	Delete(*domain.Session) error
	DeleteAll() error
}
