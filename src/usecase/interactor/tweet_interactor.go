package interactor

import (
	"github.com/IkezawaYuki/pictweet-api/src/domain/model"
	"github.com/IkezawaYuki/pictweet-api/src/domain/repository"
	"github.com/IkezawaYuki/pictweet-api/src/domain/service"
	"github.com/IkezawaYuki/pictweet-api/src/usecase/inputport"
	"time"
)

type tweetInteractor struct {
	tweetRepository repository.TweetRepository
}

func NewTweetInteractor(tweetRepo repository.TweetRepository) inputport.TweetInputport {
	return &tweetInteractor{
		tweetRepository: tweetRepo,
	}
}

func (t *tweetInteractor) FindByUUID(uuid string) (*model.Tweet, error) {
	tweet, err := t.tweetRepository.FindByUUID(uuid)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (t *tweetInteractor) Index() ([]model.Tweet, error) {
	tweets, err := t.tweetRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return tweets, nil
}

func (t *tweetInteractor) Create(userID int, text string, image string) (*model.Tweet, error) {
	tweetObj := &model.Tweet{
		Uuid:      service.CreateUUID(),
		UserID:    userID,
		Text:      text,
		Image:     image,
		CreatedAt: time.Now(),
	}
	tweet, err := t.tweetRepository.Create(tweetObj)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (t *tweetInteractor) Update(userID int, uuid string, text string, image string) (*model.Tweet, error) {
	tweetObj := &model.Tweet{
		Uuid:      uuid,
		UserID:    userID,
		Text:      text,
		Image:     image,
		CreatedAt: time.Time{},
	}
	tweet, err := t.tweetRepository.Update(tweetObj)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (t *tweetInteractor) Delete(tweet *model.Tweet) error {
	err := t.tweetRepository.Delete(tweet)
	if err != nil {
		return err
	}
	return nil
}
