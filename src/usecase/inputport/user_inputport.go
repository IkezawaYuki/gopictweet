package inputport

import "github.com/IkezawaYuki/pictweet-api/src/domain/model"

type UserInputport interface {
	CheckSession(string) (*model.Session, error)
	FindBySession(session *model.Session) (*model.User, error)
	FindByUUID(string) (*model.User, error)
}
