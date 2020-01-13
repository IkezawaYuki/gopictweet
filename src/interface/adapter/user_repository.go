package adapter

import (
	"fmt"
	"github.com/IkezawaYuki/gopictweet/src/domain"
	"github.com/IkezawaYuki/gopictweet/src/usecase"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) usecase.UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) FindAll() (users *domain.Users, err error) {
	err = ur.db.Find(&users).Error
	if err != nil {
		fmt.Printf("sql error: %s", err.Error())
	}
	return
}

// FindByEmail emailをキーにしたユーザー情報の取得
func (ur *userRepository) FindByEmail(email string) (user *domain.User, err error) {
	err = ur.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		fmt.Printf("sql error: %s", err.Error())
	}
	return
}

func (ur *userRepository) Create(user *domain.User) (*domain.User, error) {
	if err := ur.db.Create(&user).Error; err != nil {
		fmt.Printf("sql error: %s", err.Error())
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) Update(user *domain.User) (*domain.User, error) {
	if err := ur.db.Save(&user).Error; err != nil {
		fmt.Printf("sql error: %s", err.Error())
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) Delete(user *domain.User) (err error) {
	err = ur.db.Delete(&user).Error
	if err != nil {
		fmt.Printf("sql error: %s", err.Error())
	}
	return
}

// DeleteAll ユーザー情報の全削除
func (ur *userRepository) DeleteAll() (err error) {
	err = ur.db.Exec("DELETE FROM users;").Error
	if err != nil {
		fmt.Printf("sql error: %s", err.Error())
	}
	return
}
