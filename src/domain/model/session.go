package model

import "time"

type Session struct {
	ID        int
	Uuid      string
	Email     string
	UserID    int
	CreatedAt time.Time
}
