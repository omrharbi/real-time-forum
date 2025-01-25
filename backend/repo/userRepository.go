package repo

import (
	"context"
	"database/sql"
	"fmt"
	"html"
	"strings"
	"time"

	"real-time-froum/models"
)

type UserRepository interface {
	EmailExists(ctx context.Context, email string, username string) bool
	UpdateUUIDUser(ctx context.Context, uudi string, status string, userId int64, expires time.Time) error
	InsertUser(ctx context.Context, users *models.User, password string) (sql.Result, error)
	SelectUser(ctx context.Context, log *models.Login) *models.User
	CheckAuthenticat(uuid string) (bool, time.Time, int)
	CheckUser(ctx context.Context, id int) bool
	GetUserIdWithUUID(uuid string) (string, string, string, error)
	UserConnect(user int) []models.UUID
	// UpdateStatus(status string, iduser int) error
}

type userRepositoryImpl struct {
	db *sql.DB
}

// UpdateStatus implements UserRepository.

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

// UserConnect implements UserRepository.
func (u *userRepositoryImpl) UserConnect(user int) []models.UUID {
	query := `SELECT 
            id,
            username,
            firstname,
            lastname,
            status
        FROM user
		WHERE id!=?
        ORDER BY 
            CASE 
                WHEN status = 'online' THEN 1
                ELSE 2
            END,
            username ASC`
	rows, err := u.db.Query(query, user)
	if err != nil {
		fmt.Println(err)
	}
	us := []models.UUID{}
	for rows.Next() {
		ussr := models.UUID{}
		rows.Scan(&ussr.Iduser, &ussr.Nickname, &ussr.Firstname, &ussr.Lastname, &ussr.Status)
		us = append(us, ussr)
	}
	if err != nil {
		fmt.Println("error to select user", err)
	}
	return us
}

// func (u *userRepositoryImpl) UpdateStatus(status string, iduser int) error {
// 	qury := "UPDATE user SET   status=?  WHERE id=?"
// 	_, err := u.db.Exec(qury, status, iduser)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// insertUser implements UserRepository.
func (u *userRepositoryImpl) InsertUser(ctx context.Context, users *models.User, password string) (sql.Result, error) {
	Firstname := html.EscapeString(users.Firstname)
	Lastname := html.EscapeString(users.Lastname)
	Email := strings.ToLower(html.EscapeString(users.Email))
	Password := html.EscapeString(password)
	Nickname := html.EscapeString(users.Nickname)
	Gender := html.EscapeString(users.Gender)
	stm := "INSERT INTO user (username,firstname,lastname, Age ,gender ,email,password,status) VALUES(?,?,?,?,?,?,?,?)"
	row, err := u.db.ExecContext(ctx, stm, Nickname, Firstname, Lastname, users.Age, Gender, Email, Password, "online")
	return row, err
}

// selectUser implements UserRepository.
func (u *userRepositoryImpl) SelectUser(ctx context.Context, log *models.Login) *models.User {
	user := &models.User{}
	email := strings.ToLower(log.Email)
	username := strings.ToLower(log.Nickname)

	password := strings.ToLower(log.Password)
	query := "select id,email,password, firstname ,lastname FROM user where email=? or username=?"
	err := u.db.QueryRowContext(ctx, query, email, username, password).Scan(&user.Id, &user.Email, &user.Password, &user.Firstname, &user.Lastname)
	if err != nil {
		fmt.Println("error to select user", err)
	}
	return user
}

// CheckAuthenticat implements UserRepository.
func (u *userRepositoryImpl) CheckAuthenticat(uuid string) (bool, time.Time, int) {
	stm := `SELECT 
            EXISTS (SELECT 1 FROM user WHERE UUID = ?),
            (SELECT expires  FROM user WHERE UUID = ? ) AS expires,
            (SELECT id  FROM user WHERE UUID = ? ) AS id_user; `
	var exists bool
	var expires sql.NullTime
	var id any

	err := u.db.QueryRow(stm, uuid, uuid, uuid).Scan(&exists, &expires, &id)
	if err != nil {
		fmt.Println(err, "in User Repo")
		return exists, time.Time{}, 0
	}
	if !expires.Valid {
		return exists, time.Time{}, 0
	}
	if !time.Now().Before(expires.Time) {
		return false, time.Time{}, 0
	}
	if id == nil {
		return false, time.Time{}, 0
	}
	return exists, expires.Time, int(id.(int64))
}

// CheckUser implements UserRepository.
func (u *userRepositoryImpl) CheckUser(ctx context.Context, id int) bool {
	stm := `SELECT EXISTS (SELECT 1 FROM user WHERE id =  ?)  `
	var exists bool
	err := u.db.QueryRowContext(ctx, stm, id, id).Scan(&exists)
	if err != nil {
		fmt.Println(err, " in check user")
		return false
	}
	return exists
}

// emailExists implements UserRepository.
func (u *userRepositoryImpl) EmailExists(ctx context.Context, email string, nickname string) bool {
	var exists bool

	query := "SELECT EXISTS (select email from user where email=? OR username= ?)"

	err := u.db.QueryRowContext(ctx, query, email, nickname).Scan(&exists)
	if err != nil {
		fmt.Println("Error to EXISTS this Email", err)
		return false
	}
	return exists
}

// getUserIdWithUUID implements UserRepository.
func (u *userRepositoryImpl) GetUserIdWithUUID(uuid string) (string, string, string, error) {
	stm := `SELECT id , username, firstname FROM user WHERE UUID=? `
	var id_user, nickame, firstname string
	err := u.db.QueryRow(stm, uuid).Scan(&id_user, &nickame, &firstname)
	if err != nil {
		return "", "", "", err
	}
	return id_user, nickame, firstname, nil
}

// updateUUIDUser implements UserRepository.
func (u *userRepositoryImpl) UpdateUUIDUser(ctx context.Context, uudi string, status string, userId int64, expires time.Time) error {
	stm := "UPDATE user SET UUID=?,  expires =? ,status=? WHERE id=?"
	_, err := u.db.ExecContext(ctx, stm, uudi, expires, status, userId)
	return err
}
