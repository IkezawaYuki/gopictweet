package usecase

import "github.com/IkezawaYuki/gopictweet/src/domain"

type CommentRepository interface {
	FindAll() (*domain.Comments, error)
	FindByUserID(int) (*domain.Comments, error)
	FindByTweet(int) (*domain.Comments, error)
	Upsert(*domain.Comment) (*domain.Comment, error)
	Delete(*domain.Comment) error
}
