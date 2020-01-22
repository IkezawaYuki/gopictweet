package repository

import (
	"github.com/IkezawaYuki/gopictweet/src/domain/model"
)

type TweetRepository interface {
	FindAll() ([]model.Tweet, error)
	CountNumComment(int) (int, error)
	Upsert(*model.Tweet) (*model.Tweet, error)
	Delete(*model.Tweet) error
	FindByUserID(int) (*model.Tweet, error)
	FindByUUID(string) (*model.Tweet, error)
}
