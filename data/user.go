package data

import (
	"fmt"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Nickname  string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UseId     int
	CreatedAt time.Time
}

func (user *User) CreateSession() (session Session, err error) {
	statement := "insert into sessions (uuid, email, user_id, created_at) values($1, $2, $3, $4) returning id, email, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println("err is occured")
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.UseId, &session.CreatedAt)
	return
}

func (user *User) Session() (session Session, err error) {
	statement := "select id, uuid, email, user_id, created_at from sessions where user_id = $1"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	err = stmt.QueryRow().Scan(&session.Id, &session.Uuid, &session.Email, &session.UseId, &session.CreatedAt)
	return
}

func (session *Session) Check(valid bool, err error) {
	statement := "select id, uuid, email, user_id, created_at from users where uuid = $1"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()

	err = stmt.QueryRow(session.Uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UseId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

func (user *User) Delete(err error) {
	statement := "Delete from user where id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	return
}

func (user *User) Update(err error) {
	statement := "update users set nickname = $2, email = $3 where id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Nickname, user.Email)
	if err != nil {
		return
	}
}
