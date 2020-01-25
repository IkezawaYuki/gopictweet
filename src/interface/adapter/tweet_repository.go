package adapter

import (
	"fmt"
	"github.com/IkezawaYuki/gopictweet/src/domain/model"
	"github.com/IkezawaYuki/gopictweet/src/domain/repository"
	"github.com/jinzhu/gorm"
)

type tweetRepository struct {
	db *gorm.DB
}

func NewTweetRepository(db *gorm.DB) repository.TweetRepository {
	return &tweetRepository{db}
}

// FindAll tweet全取得
func (tr *tweetRepository) FindAll() (tweets []model.Tweet, err error) {
	err = tr.db.Find(&tweets).Error
	if err != nil {
		fmt.Printf("sql error: %v", err.Error())
	}
	return
}

// CountNumComment
func (tr *tweetRepository) CountNumComment(tweetID int) (num int, err error) {
	var comments []model.Comment
	err = tr.db.Where("tweet_id == ?", tweetID).Find(&comments).Error
	if err != nil {
		fmt.Printf("sql error: %v", err.Error())
	}
	// todo check
	num = len(comments)
	return
}

func (tr *tweetRepository) Create(t *model.Tweet) (*model.Tweet, error) {
	err := tr.db.Create(t).Error
	//if err != nil {
	//	fmt.Printf("sql error: %v", err.Error())
	//	return nil, err
	//}
	//return &tweet, nil
	return nil, err
}

func (tr *tweetRepository) Update(t *model.Tweet) (*model.Tweet, error) {
	err := tr.db.Update(t).Error
	//if err != nil {
	//	fmt.Printf("sql error: %v", err.Error())
	//	return nil, err
	//}
	//return &tweet, nil
	return nil, err
}

// Delete tweetの削除
func (tr *tweetRepository) Delete(tweet *model.Tweet) error {
	err := tr.db.Delete(&tweet).Error
	if err != nil {
		fmt.Printf("sql error: %v", err.Error())
	}
	return err
}

// FindByUserID ユーザーごとにtweetを取得
func (tr *tweetRepository) FindByUserID(userID int) (tweet *model.Tweet, err error) {
	err = tr.db.Where("user_id = ?", userID).Find(&tweet).Error
	if err != nil {
		fmt.Printf("sql error: %v", err.Error())
	}
	return
}

func (tr *tweetRepository) FindByUUID(uuid string) (tweet *model.Tweet, err error) {
	err = tr.db.Where("uuid = ?", uuid).Find(tweet).Error
	if err != nil {
		fmt.Printf("sql error: %v", err.Error())
	}
	return
}
