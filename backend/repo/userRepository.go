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
	CheckAuthenticat(ctx context.Context, uuid string) (bool, time.Time)
	CheckUser(ctx context.Context, id int) bool
	GetUserIdWithUUID(ctx context.Context, uuid string) (string, error)
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

// insertUser implements UserRepository.
func (u *userRepositoryImpl) InsertUser(ctx context.Context, users *models.User, password string) (sql.Result, error) {
	Firstname := html.EscapeString(users.Firstname)
	Lastname := html.EscapeString(users.Lastname)
	Email := strings.ToLower(html.EscapeString(users.Email))
	Password := html.EscapeString(password)
	Nickname := html.EscapeString(users.Nickname)
	Gender := html.EscapeString(users.Gender)
	stm := "INSERT INTO user (nickname,firstname,lastname, Age ,gender ,email,password,status) VALUES(?,?,?,?,?,?,?,?)"
	row, err := u.db.ExecContext(ctx, stm, Nickname, Firstname, Lastname, users.Age, Gender, Email, Password, "online")
	return row, err
}

// selectUser implements UserRepository.
func (u *userRepositoryImpl) SelectUser(ctx context.Context, log *models.Login) *models.User {
	user := &models.User{}
	email := strings.ToLower(log.Email)
	username := strings.ToLower(log.Nickname)

	password := strings.ToLower(log.Password)
	query := "select id,email,password, firstname ,lastname FROM user where email=? or nickname=?"
	err := u.db.QueryRowContext(ctx, query, email, username, password).Scan(&user.Id, &user.Email, &user.Password, &user.Firstname, &user.Lastname)
	if err != nil {
		fmt.Println("error to select user", err)
	}
	return user
}

// CheckAuthenticat implements UserRepository.
func (u *userRepositoryImpl) CheckAuthenticat(ctx context.Context, uuid string) (bool, time.Time) {
	stm := `SELECT 
			EXISTS (SELECT 1 FROM user WHERE UUID = ?),
			(SELECT expires FROM user WHERE UUID = ? ) AS expires; `
	var exists bool
	var expires sql.NullTime

	err := u.db.QueryRowContext(ctx, stm, uuid, uuid).Scan(&exists, &expires)
	if err != nil {
		fmt.Println(err, "in User Repo")
		return exists, time.Time{}
	}
	if !expires.Valid {
		return exists, time.Time{}
	}
	if !time.Now().Before(expires.Time) {
		return false, time.Time{}
	}
	return exists, expires.Time
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

	query := "SELECT EXISTS (select email from user where email=? OR nickname= ?)"

	err := u.db.QueryRowContext(ctx, query, email, nickname).Scan(&exists)
	if err != nil {
		fmt.Println("Error to EXISTS this Email", err)
		return false
	}
	return exists
}

// getUserIdWithUUID implements UserRepository.
func (u *userRepositoryImpl) GetUserIdWithUUID(ctx context.Context, uuid string) (string, error) {
	stm := `SELECT id FROM user WHERE UUID=? `
	var uuiduser string
	err := u.db.QueryRowContext(ctx, stm, uuid).Scan(&uuiduser)
	if err != nil {
		return "", err
	}
	return uuiduser, nil
}

// updateUUIDUser implements UserRepository.
func (u *userRepositoryImpl) UpdateUUIDUser(ctx context.Context, uudi string, status string, userId int64, expires time.Time) error {
	stm := "UPDATE user SET UUID=?,  expires =? ,status=? WHERE id=?"
	_, err := u.db.ExecContext(ctx, stm, uudi, expires, status, userId)
	return err
}
