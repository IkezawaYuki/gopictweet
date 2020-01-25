package model

import "time"

// gorm のユニークkey をUuidにせよ
type Tweet struct {
	ID        int       `gorm:"id"`
	Uuid      string    `gorm:"uuid"`
	UserID    int       `gorm:"user_id"`
	Text      string    `gorm:"text"`
	Image     string    `gorm:"image"`
	CreatedAt time.Time `gorm:"created_at"`
}
