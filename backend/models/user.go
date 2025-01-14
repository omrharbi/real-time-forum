package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type User struct {
	Id        int64     `json:"id"`
	Nickname  string    `json:"nickname"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Gender    string    `json:"gender"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
	UUID      uuid.UUID `json:"uuid"`
}
type ResponceUser struct {
	Id        int64  `json:"id"`
	Nickname  string `json:"nickname"`
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
	Nickname string `json:"nickname"`
	UUID     string `json:"uuid"`
	Password string `json:"password"`
}
type UUID struct {
	Iduser int
}
type Logout struct {
	Id   int64  `json:"id"`
	Uuid string `json:"uuid"`
}
