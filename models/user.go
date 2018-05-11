package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// User defines basic user structure
type User struct {
	ID           uint   `json:"user_id" yaml:"user_id"`
	Name         string `json:"name" yaml:"name"`
	EmailAddress string `json:"email" yaml:"email" gorm:"column:email;unique;not null"`
	LoginInfo
}

// LoginInfo defines structure used for authentication
type LoginInfo struct {
	Username     string `json:"username" yaml:"username" gorm:"unique;not null"`
	PasswordHash string `json:"-" yaml:"-" gorm:"column:password;not null"`
}

// HarpistType returns the HarpistType
func (u User) HarpistType() HarpistType {
	return USER
}

// Identity returns the ID
func (u User) Identity() uint {
	return u.ID
}

// MemberName returns the Name
func (u User) MemberName() string {
	return u.Name
}

// MemberIdentity returns the ID field used for group membership
func (u User) MemberIdentity() uint {
	return u.ID
}

// SetPassword Hashes the password string using bcrypt
func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("empty password cannot be used")
	}
	pwBytes := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	u.LoginInfo.PasswordHash = string(passwordHash)
	return nil
}

// CheckPassword validates the password string matches the bcrypt hashed password
// Example
//  if err := user.CheckPassword("passwordstring"); err != nil { failed authentication }
func (u *User) CheckPassword(password string) error {
	pwdBytes := []byte(password)
	pwdByteHashed := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(pwdByteHashed, pwdBytes)
}

// UserName returns the username of the User object
func (u User) Username() string {
	return u.LoginInfo.Username
}
