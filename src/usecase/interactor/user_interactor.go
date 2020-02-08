package interactor

import (
	"github.com/IkezawaYuki/pictweet-api/src/domain/model"
	"github.com/IkezawaYuki/pictweet-api/src/domain/repository"
	"github.com/IkezawaYuki/pictweet-api/src/usecase/inputport"
)

type userInteractor struct {
	sessionRepository repository.SessionRepository
	userRepository    repository.UserRepository
}

func NewUserInteractor(sesRepo repository.SessionRepository, userRepo repository.UserRepository) inputport.UserInputport {
	return &userInteractor{
		sessionRepository: sesRepo,
		userRepository:    userRepo,
	}
}

func (pi *userInteractor) CheckSession(uuid string) (*model.Session, error) {
	session, err := pi.sessionRepository.FindByUserID(uuid)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (pi *userInteractor) FindBySession(session *model.Session) (*model.User, error) {
	sessionID := session.Uuid
	return pi.userRepository.FindBySessionID(sessionID)
}

func (pi *userInteractor) FindByUUID(uuid string) (*model.User, error) {
	return pi.userRepository.FindByUUID(uuid)
}
