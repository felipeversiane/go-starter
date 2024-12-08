package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	id        string
	email     string
	firstName string
	lastName  string
	password  string
	search    string
	createdAt time.Time
	updatedAt time.Time
}

type UserInterface interface {
	GetID() string
	GetEmail() string
	SetEmail(email string)
	GetFirstName() string
	SetFirstName(firstName string)
	GetLastName() string
	SetLastName(lastName string)
	GetPassword() string
	SetPassword(password string) error
	GetSearch() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
}

func NewUser(email, firstName, lastName, password string) (UserInterface, error) {
	id := uuid.NewString()

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	search := generateSearch(email, firstName, lastName)

	return &user{
		id:        id,
		email:     email,
		firstName: firstName,
		lastName:  lastName,
		password:  hashedPassword,
		search:    search,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func (u *user) GetID() string {
	return u.id
}

func (u *user) GetEmail() string {
	return u.email
}

func (u *user) SetEmail(email string) {
	u.email = email
	u.updateSearch()
}

func (u *user) GetFirstName() string {
	return u.firstName
}

func (u *user) SetFirstName(firstName string) {
	u.firstName = firstName
	u.updateSearch()
}

func (u *user) GetLastName() string {
	return u.lastName
}

func (u *user) SetLastName(lastName string) {
	u.lastName = lastName
	u.updateSearch()
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

func (u *user) GetSearch() string {
	return u.search
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

func (u *user) updateSearch() {
	u.search = generateSearch(u.email, u.firstName, u.lastName)
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func generateSearch(email, firstName, lastName string) string {
	return strings.ToLower(fmt.Sprintf("%s %s %s", email, firstName, lastName))
}
