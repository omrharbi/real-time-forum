package services

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"real-time-froum/messages"
	"real-time-froum/models"
	"real-time-froum/repo"

	"github.com/gofrs/uuid/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, timeex time.Time, users *models.User) (models.ResponceUser, messages.Messages, string)
	validateUser(users *models.User) messages.Messages
	Authentication(ctx context.Context, time time.Time, log *models.Login) (models.ResponceUser, messages.Messages, uuid.UUID)
	Getuuid(ctx context.Context, uuid string)
	LogOut(ctx context.Context, uuid models.UUID) (m messages.Messages)
	checkPasswordHash(hash, password string) bool
	hashPassword(password string) string
	AuthenticatLogin(ctx context.Context, UUID string) (m messages.Messages, expire time.Time)
	UUiduser(ctx context.Context, uuid string) (m messages.Messages, us models.UUID)
	generatUUID(ctx context.Context) string
	CheckAuth(ctx context.Context, uuid string) (bool, time.Time)
}

type userServiceImpl struct {
	userRepo repo.UserRepository
}

// CheckAuthenticat implements UserService.

func NewUserService(repo repo.UserRepository) UserService {
	return &userServiceImpl{userRepo: repo}
}

func (u *userServiceImpl) CheckAuth(ctx context.Context, uuid string) (bool, time.Time) {
	return u.userRepo.CheckAuthenticat(ctx, uuid)
}

// Getuuid implements UserService.
func (u *userServiceImpl) Getuuid(ctx context.Context, uuid string) {
	panic("unimplemented")
}

// generatUUID implements UserService.
func (u *userServiceImpl) generatUUID(ctx context.Context) string {
	panic("unimplemented")
}

// NewUserService creates a new UserService

// AuthenticatLogin implements UserService.
func (u *userServiceImpl) AuthenticatLogin(ctx context.Context, UUID string) (m messages.Messages, expire time.Time) {
	exists, expire := u.userRepo.CheckAuthenticat(ctx, UUID)
	if !exists {
		m.MessageError = "Unauthorized token"
	}

	return m, expire
}

// Authentication implements UserService.
func (u *userServiceImpl) Authentication(ctx context.Context, time time.Time, log *models.Login) (models.ResponceUser, messages.Messages, uuid.UUID) {
	message := messages.Messages{}
	email := strings.ToLower(log.Email)
	username := strings.ToLower(log.Nickname)

	if (log.Nickname == "" && log.Email == "") || !u.userRepo.EmailExists(ctx, email, username) {
		message.MessageError = "Invalid email or Username"
		return models.ResponceUser{}, message, uuid.UUID{}
	} else {

		user := u.userRepo.SelectUser(ctx, log)
		if u.checkPasswordHash(user.Password, log.Password) {
			uuid, err := uuid.NewV4()
			if err != nil {
				fmt.Println("Error to Generate uuid", err)
			}
			loged := models.ResponceUser{
				Id:        user.Id,
				UUID:      uuid.String(),
				Nickname:  user.Nickname,
				Age:       user.Age,
				Gender:    user.Gender,
				Email:     user.Email,
				Firstname: user.Firstname,
				Lastname:  user.Lastname,
			}
			err = u.userRepo.UpdateUUIDUser(ctx, uuid.String(), user.Id, time)
			if err != nil {
				fmt.Println("Error to Update")
			}

			return loged, messages.Messages{}, uuid
		} else {
			message.MessageError = "Email or password incorrect."
			return models.ResponceUser{}, message, uuid.UUID{}
		}
	}
}

// LogOut implements UserService.
func (u *userServiceImpl) LogOut(ctx context.Context, uuid models.UUID) (m messages.Messages) {
	timeex := time.Now().Add(0 * time.Second)

	err := u.userRepo.UpdateUUIDUser(ctx, "null", int64(uuid.Iduser), timeex)
	if err != nil {
		m.MessageError = "Error To Update user"
		return m
	} else {
		m.MessageSucc = "Update Seccesfly"
		return m
	}
}

// UUiduser implements UserService.
func (u *userServiceImpl) UUiduser(ctx context.Context, uuid string) (m messages.Messages, us models.UUID) {
	id, err := u.userRepo.GetUserIdWithUUID(ctx, uuid)
	if err != nil {
		m.MessageError = "Unauthorized token"
		return m, models.UUID{}
	}
	id_user, err := strconv.Atoi(id)
	if err != nil {
		m.MessageError = "Unauthorized token"
		return m, models.UUID{}
	}
	us.Iduser = id_user

	return m, us
}

// checkPasswordHash implements UserService.
func (u *userServiceImpl) checkPasswordHash(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// hashPassword implements UserService.
func (u *userServiceImpl) hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("error", err)
	}
	return string(hashedPassword)
}

func generatUUID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Error to Generate uuid", err)
	}
	return uuid.String()
}

// register implements UserService.
func (u *userServiceImpl) Register(ctx context.Context, timeex time.Time, users *models.User) (models.ResponceUser, messages.Messages, string) {
	message := messages.Messages{}
	uuid := generatUUID()
	loged := models.ResponceUser{
		Id:        users.Id,
		UUID:      uuid,
		Nickname:  users.Nickname,
		Age:       users.Age,
		Gender:    users.Gender,
		Email:     users.Email,
		Firstname: users.Firstname,
		Lastname:  users.Lastname,
	}
	if strings.Trim(users.Firstname, " ") == "" ||
		strings.Trim(users.Email, " ") == "" ||
		strings.Trim(users.Lastname, " ") == "" ||
		strings.Trim(users.Password, " ") == "" ||
		strings.Trim(users.Nickname, " ") == "" ||
		strings.Trim(users.Gender, " ") == "" {
		message.MessageError = "All Input is Required"
		return models.ResponceUser{}, message, ""
	}

	message = u.validateUser(users)
	if message.MessageError != "" {
		return models.ResponceUser{}, message, ""
	}

	checkemail := strings.ToLower(users.Email)
	checkusername := strings.ToLower(users.Email)
	exists := u.userRepo.EmailExists(ctx, checkemail, checkusername)
	if exists {
		message.MessageError = "Email user already exists"
		return models.ResponceUser{}, message, ""
	}

	password := u.hashPassword(users.Password)

	rows, err := u.userRepo.InsertUser(ctx, users, password)
	if err != nil {
		message.MessageError = "Error creating this user."
		return loged, message, uuid
	}

	user_id, err := rows.LastInsertId()
	if err != nil {
		message.MessageError = err.Error()
		return models.ResponceUser{}, message, ""
	} else {
		err = u.userRepo.UpdateUUIDUser(ctx, uuid, user_id, timeex)
		if err != nil {
			fmt.Println("Error to Update")
		}
		message.MessageSucc = "User Created Successfully."
	}
	loged.Id = user_id
	return loged, message, uuid
}

// validateUser implements UserService.
func (u *userServiceImpl) validateUser(users *models.User) messages.Messages {
	message := messages.Messages{}
	nameRegex := regexp.MustCompile(`^[a-z]{2,}$`)
	if !nameRegex.MatchString(strings.TrimSpace(users.Firstname)) {
		message.MessageError = "Invalid First name"
		return message
	}

	if !nameRegex.MatchString(strings.TrimSpace(users.Lastname)) {
		message.MessageError = "Invalid Last name"
		return message
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(strings.ToLower(users.Email)) {
		message.MessageError = "Invalid email format"
		return message
	}

	if len(users.Password) < 8 {
		message.MessageError = "Invalis password length less than 8"
		return message
	}

	return message
}
