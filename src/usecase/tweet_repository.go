package usecase

import "github.com/IkezawaYuki/gopictweet/src/domain"

type TweetRepository interface {
	FindAll() (*domain.Tweets, error)
	CountNumComment(int) (int, error)
	Upsert(*domain.Tweet) (*domain.Tweet, error)
	Delete(*domain.Tweet) error
	FindByUserID(int) (*domain.Tweet, error)
}
