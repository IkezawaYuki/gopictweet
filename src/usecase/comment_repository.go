package usecase

import "github.com/IkezawaYuki/gopictweet/src/domain"

type CommentRepository interface {
	FindByUserID(int) (*domain.Comments, error)
}
