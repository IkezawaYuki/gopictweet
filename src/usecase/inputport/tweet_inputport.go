package inputport

import "github.com/IkezawaYuki/pictweet-api/src/domain/model"

type TweetInputport interface {
	Index() ([]model.Tweet, error)
	FindByUUID(string) (*model.Tweet, error)
	Create(int, string, string) (*model.Tweet, error)
	Update(int, string, string, string) (*model.Tweet, error)
	Delete(*model.Tweet) error
}
