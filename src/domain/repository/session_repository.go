package repository

import (
	"github.com/IkezawaYuki/pictweet-api/src/domain/model"
)

type SessionRepository interface {
	Create(*model.Session) (*model.Session, error)
	FindByUserID(string) (*model.Session, error)
	Check(string) (bool, error)
	Delete(*model.Session) error
	DeleteAll() error
}
