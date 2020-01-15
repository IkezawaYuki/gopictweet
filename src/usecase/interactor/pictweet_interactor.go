package interactor

import (
	"github.com/IkezawaYuki/gopictweet/src/domain"
	"github.com/IkezawaYuki/gopictweet/src/usecase"
)

type PictweetInteractor interface {
	CheckSession(string) (*domain.Session, error)
}

type pictweetInteractor struct {
	sessionRepository usecase.SessionRepository
	userRepository    usecase.UserRepository
}

func (pi *pictweetInteractor) CheckSession(uuid string) (*domain.Session, error) {
	session, err := pi.sessionRepository.FindByUserID(uuid)
	if err != nil {
		return nil, err
	}
	return session, nil
}
