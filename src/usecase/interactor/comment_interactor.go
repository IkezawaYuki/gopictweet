package interactor

import (
	"github.com/IkezawaYuki/gopictweet/src/domain"
	"github.com/IkezawaYuki/gopictweet/src/usecase"
)

type CommentInteractor interface {
	Index() (*domain.Comments, error)
	Create(*domain.Comment) (*domain.Comment, error)
	Update(*domain.Comment) (*domain.Comment, error)
	Delete(*domain.Comment) error
}

type commentInteractor struct {
	commentRepository usecase.CommentRepository
}

func NewCommentInteractor(commentRepo usecase.CommentRepository) CommentInteractor {
	return &commentInteractor{commentRepository: commentRepo}
}

func (c *commentInteractor) Index() (*domain.Comments, error) {
	comments, err := c.commentRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return comments, err
}

func (c *commentInteractor) Create(comment *domain.Comment) (*domain.Comment, error) {
	comment, err := c.commentRepository.Upsert(comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (c *commentInteractor) Update(comment *domain.Comment) (*domain.Comment, error) {
	comment, err := c.commentRepository.Upsert(comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (c *commentInteractor) Delete(comment *domain.Comment) error {
	err := c.commentRepository.Delete(comment)
	if err != nil {
		return err
	}
	return nil
}
