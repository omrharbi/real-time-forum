package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type User struct {
	Id        int64     `json:"id"`
	Nickname  string    `json:"username"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Gender    string    `json:"gender"`
	Password  string    `json:"password"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdat"`
	UUID      uuid.UUID `json:"uuid"`
}
type ResponceUser struct {
	Id        int64  `json:"id"`
	Nickname  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
	UUID      string `json:"uuid"`
}
type Login struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"username"`
	UUID     string `json:"uuid"`
	Password string `json:"password"`
}
type UUID struct {
	Iduser    int    `json:"id"`
	Nickname  string `json:"username"`
	Firstname string `json:"firstname"`
	Status    string `json:"status"`
	Lastname  string `json:"lastname"`
	Seen      any    `json:"seen"`
}
type Logout struct {
	Id   int    `json:"id"`
	Uuid string `json:"uuid"`
}
