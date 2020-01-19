package inputport

import "github.com/IkezawaYuki/gopictweet/src/domain/model"

type UserInputport interface {
	CheckSession(string) (*model.Session, error)
	FindUserBySession(session *model.Session) (*model.User, error)
}
