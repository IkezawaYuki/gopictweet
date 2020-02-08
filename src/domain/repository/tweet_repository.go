package repository

import (
	"github.com/IkezawaYuki/pictweet-api/src/domain/model"
)

type TweetRepository interface {
	FindAll() ([]model.Tweet, error)
	CountNumComment(int) (int, error)
	Create(*model.Tweet) (*model.Tweet, error)
	Update(*model.Tweet) (*model.Tweet, error)
	Delete(*model.Tweet) error
	FindByUserID(int) (*model.Tweet, error)
	FindByUUID(string) (*model.Tweet, error)
}
