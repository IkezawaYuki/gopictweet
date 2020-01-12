package domain

import "time"

type Tweet struct {
	ID        int
	UuID      string
	UserID    int
	Text      string
	Image     string
	CreatedAt time.Time
}

type Tweets []Tweet