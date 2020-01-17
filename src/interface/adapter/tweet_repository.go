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

// FindAll tweet全取得
func (tr *tweetRepository) FindAll() (tweets *domain.Tweets, err error) {
	err = tr.db.Find(&tweets).Error
	if err != nil {
		fmt.Printf("sql error: %v", err.Error())
	}
	return
}

// CountNumComment
func (tr *tweetRepository) CountNumComment(tweetID int) (num int, err error) {
	var comments []domain.Comment
	err = tr.db.Where("tweet_id == ?", tweetID).Find(&comments).Error
	if err != nil {
		fmt.Printf("sql error: %v", err.Error())
	}
	// todo check
	num = len(comments)
	return
}

// Upsert 同じプライマリーキーを持つ物を見つけたらupdate見つからない場合はinsert
func (tr *tweetRepository) Upsert(tweet *domain.Tweet) (result *domain.Tweet, err error) {
	err = tr.db.Where(domain.Tweet{ID: tweet.ID}).Attrs(domain.Tweet{
		UuID:   tweet.UuID,
		UserID: tweet.UserID,
		Text:   tweet.Text,
		Image:  tweet.Image,
	}).FirstOrCreate(&tweet).Scan(&result).Error
	if err != nil {
		fmt.Printf("sql error: %v", err.Error())
	}
	return
}

// Delete tweetの削除
func (tr *tweetRepository) Delete(tweet *domain.Tweet) error {
	err := tr.db.Delete(&tweet).Error
	if err != nil {
		fmt.Printf("sql error: %v", err.Error())
	}
	return err
}

// FindByUserID ユーザーごとにtweetを取得
func (tr *tweetRepository) FindByUserID(userID int) (tweet *domain.Tweet, err error) {
	err = tr.db.Where("user_id = ?", userID).Find(&tweet).Error
	if err != nil {
		fmt.Printf("sql error: %v", err.Error())
	}
	return
}

func (tr *tweetRepository) FindByUUID(uuid string) (tweet *domain.Tweet, err error) {
	err = tr.db.Where("uuid = ?", uuid).Find(tweet).Error
	if err != nil {
		fmt.Printf("sql error: %v", err.Error())
	}
	return
}
