package models

import (
	"github.com/go-harpist/harpist"
	"testing"
)

var TestUser = User{
	Name:         "Test User",
	EmailAddress: "test@example.com",
	LoginInfo:    LoginInfo{Username: "test_user"},
}

func TestUserCreate(t *testing.T) {
	var err error

	harpist.HarpistDB.Where(TestUser).First(&TestUser)

	if harpist.HarpistDB.NewRecord(TestUser) {
		TestUser.SetPassword("password")
		harpist.HarpistDB.Create(&TestUser)
		err = harpist.HarpistDB.Save(&TestUser).Error
	}

	if err != nil {
		t.Error("Unable to store user in the database")
	}

	user2 := User{}

	harpist.HarpistDB.Where(User{Name: "Test User"}).First(&user2)

	if user2.LoginInfo.Username != "test_user" {
		t.Error("User does not match or does not fetch from database")
	}
}

func TestUserHarpistType(t *testing.T) {
	user := User{}

	harpist.HarpistDB.Where(TestUser).First(&user)
	if user.HarpistType() != USER {
		t.Error("HarpistType() does not return USER")
	}
}

func TestUserIdentity(t *testing.T) {
	user := User{}
	harpist.HarpistDB.Where(TestUser).First(&user)
	if user.Identity() != 1 {
		t.Error("Identity() does not test user ID 1")
	}
	if user.MemberIdentity() != 1 {
		t.Error("MemberIdentity() does not test user ID 1")
	}
}

func TestUserMemberName(t *testing.T) {
	user := User{}

	harpist.HarpistDB.Where(TestUser).First(&user)
	if user.MemberName() != "Test User" {
		t.Error("MemberName() does not return 'Test User'")
	}
}

func TestUserSetAndCheckPassword(t *testing.T) {
	user := User{}

	harpist.HarpistDB.Where(TestUser).First(&user)
	err := user.SetPassword("password1")
	if err != nil {
		t.Errorf("Error encrypting password: %v", err)
	}
	harpist.HarpistDB.Where(TestUser).First(&user)
	err = user.CheckPassword("password")
	if err != nil {
		t.Error("Error checking password against test database")
	}
}

func TestUserUsername(t *testing.T) {
	user := User{}

	harpist.HarpistDB.Where(TestUser).First(&user)
	if user.Username() != "test_user" {
		t.Error("Got incorrect username for test_user")
	}
}
