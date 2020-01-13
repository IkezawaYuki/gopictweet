package domain

import "time"

type Comment struct {
	ID        int
	UuID      string
	UserID    int
	TweetID   int
	Text      string
	CreatedAt time.Time
}

type Comments []Comment
