package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type User struct {
	Id        int64     `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
	UUID      uuid.UUID `json:"uuid"`
}
type ResponceUser struct {
	Id        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	UUID      string `json:"uuid"`
}
type Login struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
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
