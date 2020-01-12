package adapter

import (
	"fmt"
	"github.com/IkezawaYuki/gopictweet/src/domain"
	"github.com/IkezawaYuki/gopictweet/src/usecase"
	"github.com/jinzhu/gorm"
)

type tweetRepository struct {
	db *gorm.DB
}

func NewTweetRepository(db *gorm.DB) usecase.TweetRepository {
	return &tweetRepository{db}
}

func (tr *tweetRepository) FindAll() (tweets *domain.Tweets, err error) {
	err = tr.db.Find(&tweets).Error
	if err != nil {
		fmt.Printf("sql error: %v", err.Error())
	}
	return
}

func (tr *tweetRepository) CountNumComment() (int, error) {

}
func (tr *tweetRepository) Create(*domain.Tweet) (*domain.Tweet, error) {

}
func (tr *tweetRepository) Update(*domain.Tweet) (*domain.Tweet, error) {

}
func (tr *tweetRepository) Delete(*domain.Tweet) error {

}
func (tr *tweetRepository) FindByUserID(int) (*domain.Tweet, error) {

}
