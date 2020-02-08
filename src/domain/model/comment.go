package model

import "time"

type Comment struct {
	ID        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	UserID    int       `json:"user_id"`
	TweetID   int       `json:"tweet_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type Comments []Comment
