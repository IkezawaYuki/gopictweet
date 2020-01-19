package interactor

import (
	"github.com/IkezawaYuki/gopictweet/src/domain/model"
	"github.com/IkezawaYuki/gopictweet/src/domain/repository"
	"github.com/IkezawaYuki/gopictweet/src/usecase/inputport"
	"time"
)

type commentInteractor struct {
	commentRepository repository.CommentRepository
}

func NewCommentInteractor(commentRepo repository.CommentRepository) inputport.CommentInputport {
	return &commentInteractor{commentRepository: commentRepo}
}

func (c *commentInteractor) Index() (*model.Comments, error) {
	comments, err := c.commentRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return comments, err
}

func (c *commentInteractor) Create(userID int, tweetID int, text string) (*model.Comment, error) {
	comment := &model.Comment{
		UuID:      "", // todo createUUID
		UserID:    userID,
		TweetID:   tweetID,
		Text:      text,
		CreatedAt: time.Now(),
	}
	comment, err := c.commentRepository.Upsert(comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (c *commentInteractor) Update(comment *model.Comment) (*model.Comment, error) {
	comment, err := c.commentRepository.Upsert(comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (c *commentInteractor) Delete(comment *model.Comment) error {
	err := c.commentRepository.Delete(comment)
	if err != nil {
		return err
	}
	return nil
}

func (c *commentInteractor) FindByTweetID(id int) (*model.Comments, error) {
	comments, err := c.commentRepository.FindByTweet(id)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
