package interactor

import (
	"github.com/IkezawaYuki/gopictweet/src/domain"
	"github.com/IkezawaYuki/gopictweet/src/usecase"
	"time"
)

type TweetInteractor interface {
	Index() (*domain.Tweets, error)
	FindByUUID(string) (*domain.Tweet, error)
	Create(int, string, string) (*domain.Tweet, error)
	Update(int, string, string, string) (*domain.Tweet, error)
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

func (t *tweetInteractor) FindByUUID(uuid string) (*domain.Tweet, error) {
	tweet, err := t.tweetRepository.FindByUUID(uuid)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (t *tweetInteractor) Index() (*domain.Tweets, error) {
	tweets, err := t.tweetRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return tweets, nil
}

func (t *tweetInteractor) Create(userID int, text string, image string) (*domain.Tweet, error) {
	tweetObj := &domain.Tweet{
		UuID:      "",
		UserID:    userID,
		Text:      text,
		Image:     image,
		CreatedAt: time.Now(),
	}
	tweet, err := t.tweetRepository.Upsert(tweetObj)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (t *tweetInteractor) Update(userID int, uuid string, text string, image string) (*domain.Tweet, error) {
	tweetObj := &domain.Tweet{
		UuID:      uuid,
		UserID:    userID,
		Text:      text,
		Image:     image,
		CreatedAt: time.Time{},
	}
	tweet, err := t.tweetRepository.Upsert(tweetObj)
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
