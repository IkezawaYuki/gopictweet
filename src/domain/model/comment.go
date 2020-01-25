package model

import "time"

type Comment struct {
	ID        int
	Uuid      string
	UserID    int
	TweetID   int
	Text      string
	CreatedAt time.Time
}

type Comments []Comment
