package adapter

import (
	"fmt"
	"github.com/IkezawaYuki/pictweet-api/src/domain/model"
	"github.com/IkezawaYuki/pictweet-api/src/domain/repository"
	"github.com/jinzhu/gorm"
)

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) repository.SessionRepository {
	return &sessionRepository{db: db}
}

func (sr *sessionRepository) Create(session *model.Session) (*model.Session, error) {
	err := sr.db.Create(&session).Error
	if err != nil {
		fmt.Printf("sql error: %v", err.Error())
		return nil, err
	}
	return session, nil
}

func (sr *sessionRepository) FindByUserID(userID string) (*model.Session, error) {
	var session model.Session
	err := sr.db.Where("user_id = ?", userID).Find(&session).Error
	if err != nil {
		fmt.Printf("sql error: %v", err.Error())
		return nil, err
	}
	return &session, nil
}

func (sr *sessionRepository) Check(uuid string) (bool, error) {
	var session model.Session
	if err := sr.db.Where("uuid = ?", uuid).Find(&session).Error; err != nil {
		fmt.Printf("sql error: %v", err.Error())
		return false, nil
	}
	return true, nil
}

func (sr *sessionRepository) Delete(session *model.Session) error {
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
