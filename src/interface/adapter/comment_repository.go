package adapter

import (
	"fmt"
	"github.com/IkezawaYuki/gopictweet/src/domain/model"
	"github.com/IkezawaYuki/gopictweet/src/domain/repository"
	"github.com/jinzhu/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

// NewCommentRepository di
func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return &commentRepository{db: db}
}

func (cr *commentRepository) FindAll() (comments *model.Comments, err error) {
	err = cr.db.Find(&comments).Error
	if err != nil {
		fmt.Printf("sql error: %v", err)
	}
	return
}

func (cr *commentRepository) FindByUserID(userID int) (comments *model.Comments, err error) {
	err = cr.db.Where("user_id", userID).Find(&comments).Error
	if err != nil {
		fmt.Printf("sql error: %v", err)
	}
	return
}

func (cr *commentRepository) FindByTweet(tweetID int) (comments *model.Comments, err error) {
	err = cr.db.Where("tweet_id", tweetID).Find(&comments).Error
	if err != nil {
		fmt.Printf("sql error: %v", err)
	}
	return
}

func (cr *commentRepository) Upsert(comment *model.Comment) (result *model.Comment, err error) {
	err = cr.db.Where(model.Comment{ID: comment.ID}).Attrs(model.Comment{
		Uuid:    comment.Uuid,
		UserID:  comment.UserID,
		TweetID: comment.TweetID,
		Text:    comment.Text,
	}).FirstOrCreate(&result).Error
	if err != nil {
		fmt.Printf("sql error: %v", err)
	}
	return
}

func (cr *commentRepository) Delete(comment *model.Comment) (err error) {
	err = cr.db.Delete(&comment).Error
	if err != nil {
		fmt.Printf("sql error: %v", err)
	}
	return
}
