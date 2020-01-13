package interactor

import "github.com/IkezawaYuki/gopictweet/src/usecase"

type PictweetInteractor struct {
	sessionRepository usecase.SessionRepository
	userRepository    usecase.UserRepository
}
