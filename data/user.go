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
	statement := "insert into sessions (uuid, email, user_id, created_at) values($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UseId, &session.CreatedAt)
	return
}

func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("select id, uuid, email, user_id, created_at from users where user_id = $1", user.Uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UseId, &session.CreatedAt)
	return
}

func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("select id, uuid, email, user_id, created_at from sessions where uuid = $1").Scan(&session.Id, &session.Uuid, &session.Email, &session.UseId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

func (session *Session) DeleteByUUID() (err error) {
	statement := "delete from session where uuid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println("session is none")
	}
	_, err = stmt.Exec(session.Uuid)
	return
}

func (user *User) Create() (err error) {
	statement := "insert into users (uuid, email, password, nickname, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), user.Email, Encrypt(user.Password), user.Nickname, time.Now()).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}

func (user *User) Delete() (err error) {
	statement := "Delete from users where id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	return
}

func (user *User) Update() (err error) {
	statement := "update users set nickname = $2, email = $3 where id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Nickname, user.Email)
	return
}

func (user *User) DeleteAll() (err error) {
	statement := "delete from users"
	_, err = Db.Exec(statement)
	return
}

func Users() (users []User, err error) {
	statement := "select id, uuid, nickname, email, password, created_at from users"
	rows, err := Db.Query(statement)
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Uuid, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	return
}

func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id, uuid, nickname, email, password, created_at from users where email = $1", email).
		Scan(&user.Id, &user.Uuid, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func UserByUUID(uuid string) (user User, err error) {
	statement := "select id, uuid, nickname, email, password, created_at from users where uuid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	err = stmt.QueryRow(uuid).Scan(&user.Id, &user.Uuid, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func (session *Session) User() (user User, err error) {
	statement := "select id, uuid, nickname, email, password, created_at from users where id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	err = stmt.QueryRow(statement, session.Id).Scan(&user.Id, &user.Uuid, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func UserDeleteAll() (err error) {
	statement := "delete from users"
	_, err = Db.Exec(statement)
	return
}

func SessionDeleteAll() (err error) {
	statement := "delete from sessions"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	return
}
