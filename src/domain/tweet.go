package domain

import "time"

type Tweet struct {
	ID        int       `gorm:"id"`
	UuID      string    `gorm:"uuid"`
	UserID    int       `gorm:"user_id"`
	Text      string    `gorm:"text"`
	Image     string    `gorm:"image"`
	CreatedAt time.Time `gorm:created_at`
}

type Tweets []Tweet
