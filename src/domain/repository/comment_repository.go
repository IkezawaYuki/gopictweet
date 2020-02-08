package repository

import (
	"github.com/IkezawaYuki/pictweet-api/src/domain/model"
)

type CommentRepository interface {
	FindAll() (*model.Comments, error)
	FindByUserID(int) (*model.Comments, error)
	FindByTweet(int) (*model.Comments, error)
	Upsert(*model.Comment) (*model.Comment, error)
	Delete(*model.Comment) error
}
