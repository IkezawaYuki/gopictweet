package domain

import "time"

type User struct {
	Id        int    `gorm:"id"`
	Uuid      string `gorm:"uuid"`
	Nickname  string `gorm:"nickname"`
	Email     string `gorm:"email"`
	Password  string `gorm:"password"`
	CreatedAt time.Time
}

type Users []User
