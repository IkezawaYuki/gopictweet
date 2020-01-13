package adapter

import (
	"fmt"
	"github.com/IkezawaYuki/gopictweet/src/domain"
	"github.com/IkezawaYuki/gopictweet/src/usecase"
	"github.com/jinzhu/gorm"
)

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) usecase.SessionRepository {
	return &sessionRepository{db: db}
}

func (sr *sessionRepository) Create(session *domain.Session) (*domain.Session, error) {
	err := sr.db.Create(&session).Error
	if err != nil {
		fmt.Printf("sql error: %v", err.Error())
		return nil, err
	}
	return session, nil
}

func (sr *sessionRepository) FindByUserID(userID string) (*domain.Session, error) {
	var session domain.Session
	err := sr.db.Where("user_id = ?", userID).Find(&session).Error
	if err != nil {
		fmt.Printf("sql error: %v", err.Error())
		return nil, err
	}
	return &session, nil
}

func (sr *sessionRepository) Check(uuid string) (bool, error) {
	var session domain.Session
	if err := sr.db.Where("uuid = ?", uuid).Find(&session).Error; err != nil {
		fmt.Printf("sql error: %v", err.Error())
		return false, nil
	}
	return true, nil
}

func (sr *sessionRepository) Delete(session *domain.Session) error {
	if err := sr.db.Delete(&session).Error; err != nil {
		fmt.Printf("sql error: %v", err.Error())
		return err
	}
	return nil
}

func (sr *sessionRepository) DeleteAll() error {
	err := sr.db.Exec("DELETE FROM sessions;").Error
	if err != nil {
		fmt.Printf("sql error: %v", err.Error())
		return err
	}
	return nil
}
