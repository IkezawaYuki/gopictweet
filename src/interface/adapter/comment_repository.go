package adapter

import (
	"github.com/IkezawaYuki/gopictweet/src/domain"
	"github.com/IkezawaYuki/gopictweet/src/usecase"
	"github.com/jinzhu/gorm"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

type commentRepository struct {
	db *gorm.DB
}

// NewCommentRepository di
func NewCommentRepository(db *gorm.DB) usecase.CommentRepository {
	return &commentRepository{db: db}
}

func (cr *commentRepository) FindAll() (comments *domain.Comments, err error) {
	err = cr.db.Find(&comments).Error
	if err != nil {
		fmt.Printf("sql error: %v", err)
	}
	return
}

func (cr *commentRepository) FindByUserID(userID int) (comments *domain.Comments, err error) {
	err = cr.db.Where("user_id", userID).Find(&comments).Error
	if err != nil {
		fmt.Printf("sql error: %v", err)
	}
	return
}

func (cr *commentRepository) FindByTweet(tweetID int) (comments *domain.Comments, err error) {
	err = cr.db.Where("tweet_id", tweetID).Find(&comments).Error
	if err != nil {
		fmt.Printf("sql error: %v", err)
	}
	return
}

func (cr *commentRepository) Upsert(comment *domain.Comment) (result *domain.Comment, err error) {
	err = cr.db.Where(domain.Comment{ID: comment.ID}).Attrs(domain.Comment{
		UuID:    comment.UuID,
		UserID:  comment.UserID,
		TweetID: comment.TweetID,
		Text:    comment.Text,
	}).FirstOrCreate(&result).Error
	if err != nil {
		fmt.Printf("sql error: %v", err)
	}
	return
}

func (cr *commentRepository) Delete(comment *domain.Comment) (err error) {
	err = cr.db.Delete(&comment).Error
	if err != nil {
		fmt.Printf("sql error: %v", err)
	}
	return
}
