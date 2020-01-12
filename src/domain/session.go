package domain

import "time"

type Session struct {
	ID        int
	UuID      string
	Email     string
	UserID     int
	CreatedAt time.Time
}