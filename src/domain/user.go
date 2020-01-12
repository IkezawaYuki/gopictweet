package domain

import "time"

type User struct {
	Id        int
	Uuid      string
	Nickname  string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Users []User