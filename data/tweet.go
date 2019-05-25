package data

import (
	"fmt"
	"time"
)

type Tweet struct {
	Id        int
	Uuid      string
	UserId    int
	Text      string
	Image     string
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Uuid      string
	UserId    int
	TweetId   int
	Text      string
	CreatedAt time.Time
}

func (tweet *Tweet) CreatedAtDate() string {
	return tweet.CreatedAt.Format("2006-01-02 15:04:05")
}

func (comment *Comment) CreatedAtDate() string {
	return comment.CreatedAt.Format("2006-01-02 15:04:05")
}

func (tweet *Tweet) NumComment() (count int) {
	statement := "select count(*) from comments where tweet_id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, _ := stmt.Query(tweet.Id)
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return
		}
	}
	return
}

func (tweet *Tweet) Comments() (comments []Comment, err error) {
	statement := "select id, uuid, user_id, tweet_id, text, created_at from comments where tweet_id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(tweet.Id)
	for rows.Next() {
		var comment Comment
		err = rows.Scan(&comment.Id, &comment.Uuid, &comment.UserId, &comment.TweetId, &comment.Text, &comment.CreatedAt)
		if err != nil {
			return
		}
		comments = append(comments, comment)
	}
	return
}

func (user *User) CreateTweet(text string, image string) (tweet Tweet, err error) {
	fmt.Println("createTweet 通過")
	statement := "insert into tweets (uuid, user_id, text, image, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, user_id, text, image, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), user.Id, text, image, time.Now()).Scan(&tweet.Id, &tweet.Uuid, &tweet.UserId, &tweet.Text, &tweet.Image, &tweet.CreatedAt)
	return
}

func (user *User) CreateComment(tweet Tweet, text string) (comment Comment, err error) {
	statement := "insert into comments (uuid, user_id, tweet_id, text, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, user_id, tweet_id, text, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), user.Id, tweet.Id, text, time.Now()).Scan(&comment.Id, &comment.Uuid, &comment.UserId, &comment.TweetId, &comment.CreatedAt)
	return
}

func Tweets() (tweets []Tweet, err error) {
	rows, err := Db.Query("select id, uuid, user_id, text, image, created_at from tweets")
	if err != nil {
		return
	}
	for rows.Next() {
		tweet := Tweet{}
		err = rows.Scan(&tweet.Id, &tweet.Uuid, &tweet.UserId, &tweet.Text, &tweet.Image, &tweet.CreatedAt)
		if err != nil {
			return
		}
		tweets = append(tweets, tweet)
	}
	rows.Close()
	return
}

func TweetByUuid(uuid string) (tweet Tweet, err error) {
	statement := "select id, uuid, user_id, text, image, created_at from tweets where uuid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(uuid).Scan(&tweet.Id, &tweet.Uuid, &tweet.UserId, &tweet.Text, &tweet.Image, &tweet.CreatedAt)
	return
}

func (tweet *Tweet) User() (user User, err error) {
	user = User{}
	Db.QueryRow("select id, uuid, nickname, email, password, created_at from users where id = $1",
		tweet.UserId).Scan(&user.Id, &user.Uuid, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func (comment *Comment) User() (user User, err error) {
	user = User{}
	Db.QueryRow("select id, uuid, nickname, email, password, created_at from users where id = $1",
		comment.UserId).Scan(&user.Id, &user.Uuid, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt)
	return
}
