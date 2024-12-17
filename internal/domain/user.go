package domain

import (
	"fmt"
	"time"

	"github.com/felipeversiane/go-starter/internal/infra/config/response"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	id        string
	email     string
	firstName string
	lastName  string
	password  string
	createdAt time.Time
	updatedAt time.Time
	deleted   bool
}

type UserInterface interface {
	GetID() string
	GetEmail() string
	GetFirstName() string
	SetFirstName(firstName string)
	GetLastName() string
	SetLastName(lastName string)
	GetPassword() string
	SetPassword(password string) error
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
	GetDeleted() bool
	SetDeleted(deleted bool)
	GenerateToken() (string, string, *response.ErrorResponse)
	GenerateAcessToken() (string, *response.ErrorResponse)
	GenerateRefreshToken() (string, *response.ErrorResponse)
}

func NewUser(email, firstName, lastName, password string) (UserInterface, error) {
	id := uuid.NewString()

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	return &user{
		id:        id,
		email:     email,
		firstName: firstName,
		lastName:  lastName,
		password:  hashedPassword,
		createdAt: time.Now(),
		updatedAt: time.Now(),
		deleted:   false,
	}, nil
}

func NewUpdateUser(firstName, lastName string) UserInterface {
	user := &user{
		firstName: firstName,
		lastName:  lastName,
		updatedAt: time.Now(),
	}

	return user
}

func (u *user) GetDeleted() bool {
	return u.deleted
}

func (u *user) SetDeleted(deleted bool) {
	u.deleted = deleted
}

func (u *user) GetID() string {
	return u.id
}

func (u *user) GetEmail() string {
	return u.email
}

func (u *user) GetFirstName() string {
	return u.firstName
}

func (u *user) SetFirstName(firstName string) {
	u.firstName = firstName
}

func (u *user) GetLastName() string {
	return u.lastName
}

func (u *user) SetLastName(lastName string) {
	u.lastName = lastName
}

func (u *user) GetPassword() string {
	return u.password
}

func (u *user) SetPassword(password string) error {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}
	u.password = hashedPassword
	return nil
}

func (u *user) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *user) GetUpdatedAt() time.Time {
	return u.updatedAt
}

func (u *user) SetUpdatedAt(updatedAt time.Time) {
	u.updatedAt = updatedAt
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
