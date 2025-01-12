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
	emailExists(ctx context.Context, email string) bool
	updateUUIDUser(ctx context.Context, uudi string, userId int64, expires time.Time) error
	insertUser(ctx context.Context, users *models.User, password string) (sql.Result, error)
	selectUser(ctx context.Context, log *models.Login) *models.User
	CheckAuthenticat(ctx context.Context, uuid string) (bool, time.Time)
	CheckUser(ctx context.Context, id int) bool
	getUserIdWithUUID(ctx context.Context, uuid string) (string, error)
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

// insertUser implements UserRepository.
func (u *userRepositoryImpl) insertUser(ctx context.Context,users *models.User, password string) (sql.Result, error) {
	Firstname := html.EscapeString(users.Firstname)
	Lastname := html.EscapeString(users.Lastname)
	Email := strings.ToLower(html.EscapeString(users.Email))
	Password := html.EscapeString(password)
	stm := "INSERT INTO user (firstname,lastname,email,password) VALUES(?,?,?,?)"
	row, err := u.db.ExecContext(ctx, stm, Firstname, Lastname, Email, Password)
	return row, err
}

// selectUser implements UserRepository.
func (u *userRepositoryImpl) selectUser(ctx context.Context,log *models.Login) *models.User {
	user := models.User{}
	email := strings.ToLower(log.Email)
	password := strings.ToLower(log.Password)
	query := "select id,email,password, firstname ,lastname FROM user where email=?"
	err := u.db.QueryRowContext(ctx, query, email, password).Scan(&user.Id, &user.Email, &user.Password, &user.Firstname, &user.Lastname)
	if err != nil {
		fmt.Println("error to select user", err)
	}
	return &user
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
		fmt.Println(err, "here")
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
func (u *userRepositoryImpl) CheckUser(ctx context.Context,id int) bool {
	stm := `SELECT EXISTS (SELECT 1 FROM user WHERE id =  ?)  `
	var exists bool
	err := u.db.QueryRowContext(ctx, stm, id, id).Scan(&exists)
	if err != nil {
		fmt.Println(err)
	}
	return exists
}

// emailExists implements UserRepository.
func (u *userRepositoryImpl) emailExists(ctx context.Context,email string) bool {
	var exists bool
	query := "SELECT EXISTS (select email from user where email=?)"
	err := u.db.QueryRowContext(ctx, query, email).Scan(&exists)
	if err != nil {
		fmt.Println("Error to EXISTS this Email", err)
	}
	return exists
}

// getUserIdWithUUID implements UserRepository.
func (u *userRepositoryImpl) getUserIdWithUUID(ctx context.Context,uuid string) (string, error) {
	stm := `SELECT id FROM user WHERE UUID=? `
	var uuiduser string
	err := u.db.QueryRowContext(ctx, stm, uuid).Scan(&uuiduser)
	if err != nil {
		return "", err
	}
	return uuiduser, nil
}

// updateUUIDUser implements UserRepository.
func (u *userRepositoryImpl) updateUUIDUser(ctx context.Context,uudi string, userId int64, expires time.Time) error {
	stm := "UPDATE user SET UUID=?, expires =?  WHERE id=?"
	_, err := u.db.ExecContext(ctx,  stm, uudi, expires, userId)
	return err
}
