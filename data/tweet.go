package data

import "time"

type Tweet struct {
	Id        int
	Uuid      string
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
