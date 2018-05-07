package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// User defines basic user structure
type User struct {
	ID   uint   `json:"user_id" yaml:"user_id"`
	Name string `json:"name" yaml:"name"`
	LoginInfo
}

type LoginInfo struct {
	Username     string `json:"username" yaml:"username" gorm:"unique;not null"`
	PasswordHash string `json:"password" yaml:"password" gorm:"column:password;not null"`
}

func (u User) HarpistType() HarpistType {
	return USER
}

func (u User) Identity() uint {
	return u.ID
}

func (u User) MemberName() string {
	return u.Name
}

func (u User) MemberIdentity() uint {
	return u.ID
}

func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("Empty password cannot be used.")
	}
	pwBytes := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	u.LoginInfo.PasswordHash = string(passwordHash)
	return nil
}

// Example
//  if err := user.CheckPassword("passwordstring"); err != nil { failed authentication }
func (u *User) CheckPassword(password string) error {
	pwdBytes := []byte(password)
	pwdByteHashed := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(pwdByteHashed, pwdBytes)
}

func (u User) UserName() string {
	return u.LoginInfo.Username
}
