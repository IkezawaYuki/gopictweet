package interactor

import (
	"github.com/IkezawaYuki/gopictweet/src/domain"
	"github.com/IkezawaYuki/gopictweet/src/usecase"
)

type TweetInteractor interface {
	Index() (*domain.Tweets, error)
	Create(*domain.Tweet) (*domain.Tweet, error)
	Update(*domain.Tweet) (*domain.Tweet, error)
	Delete(*domain.Tweet) error
}

type tweetInteractor struct {
	tweetRepository usecase.TweetRepository
}

func NewTweetInteractor(tweetRepo usecase.TweetRepository) TweetInteractor {
	return &tweetInteractor{
		tweetRepository: tweetRepo,
	}
}

func (t *tweetInteractor) Index() (*domain.Tweets, error) {
	tweets, err := t.tweetRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return tweets, nil
}

func (t *tweetInteractor) Create(tweet *domain.Tweet) (*domain.Tweet, error) {
	tweet, err := t.tweetRepository.Upsert(tweet)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (t *tweetInteractor) Update(tweet *domain.Tweet) (*domain.Tweet, error) {
	tweet, err := t.tweetRepository.Upsert(tweet)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (t *tweetInteractor) Delete(tweet *domain.Tweet) error {
	err := t.tweetRepository.Delete(tweet)
	if err != nil {
		return err
	}
	return nil
}
