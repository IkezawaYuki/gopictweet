package inputport

import "github.com/IkezawaYuki/pictweet-api/src/domain/model"

type CommentInputport interface {
	Index() (*model.Comments, error)
	Create(int, int, string) (*model.Comment, error)
	Update(*model.Comment) (*model.Comment, error)
	Delete(*model.Comment) error
	FindByTweetID(int) (*model.Comments, error)
}
