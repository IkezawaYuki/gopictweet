package data

import "testing"

func TweetDeleteAll() (err error) {
	db := db()
	defer db.Close()
	statement := "delete from tweets"
	_, err := db.Exec(statement)
	if err != nil {
		return
	}
	return
}

func Test_CreateTweet(t *testing.T) {
	setup()
	err := users[0].Create()
	if err != nil {
		t.Error(err, "Cannot create user.")
	}
	conv, err := users[0].CreateComment("My first tweeet")
	if err != nil {
		t.Error(err, "Cannot create tweet")
	}
	if conv.UserId != users[0].Id {
		t.Error("User not liked with tweet")
	}
}
