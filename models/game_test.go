package models

import (
	"github.com/go-harpist/harpist"
	"testing"
)

var TestGame Game

func init() {
	TestGame = Game{
		Name: "Test Game",
	}

	harpist.HarpistDB.Where(TestUser).First(&TestUser)

	if harpist.HarpistDB.NewRecord(TestUser) {
		TestUser.SetPassword("password")
		harpist.HarpistDB.Create(&TestUser)
		_ = harpist.HarpistDB.Save(&TestUser).Error
	}

	if harpist.HarpistDB.NewRecord(TestGame) {
		TestGame.Owner = TestUser
		_ = harpist.HarpistDB.Save(&TestGame).Error
	}
}

func TestGameCreate(t *testing.T) {
	var err error

	if harpist.HarpistDB.NewRecord(TestGame) {
		TestGame.Owner = TestUser
		err = harpist.HarpistDB.Save(&TestGame).Error
	}

	if err != nil {
		t.Error("Unable to store game in the database")
	}

}

func TestGameHarpistType(t *testing.T) {
	game := Game{}

	harpist.HarpistDB.Where(TestGame).First(&game)
	if game.HarpistType() != GAME {
		t.Error("HarpistType() does not return GAME")
	}
}

func TestGameIdentity(t *testing.T) {
	game := Game{}

	harpist.HarpistDB.First(&game)
	if game.Identity() != 1 {
		t.Errorf("Identity() does not test game ID 1: %v", game.Identity())
	}
	if game.MemberIdentity() != 1 {
		t.Error("MemberIdentity() does not test game ID 1")
	}
}

func TestGameOwner(t *testing.T) {
	game := Game{}

	harpist.HarpistDB.Where(TestGame).First(&game)
	user := game.GroupOwner()
	harpist.HarpistDB.Where(user).First(&user)
	if game.GroupOwner().(User).ID != user.(User).ID {
		t.Error("Did not fetch correct user object")
	}
}

func TestGameMemberName(t *testing.T) {
	game := Game{}

	harpist.HarpistDB.First(&game)
	if game.MemberName() != "Test Game" {
		t.Errorf("MemberName() does not return 'Test Game', returns %v", game.MemberName())
	}
}

func TestGameAddMember(t *testing.T) {
	game := Game{}

	harpist.HarpistDB.First(&game)
	game.AddMember(TestUser)
	err := harpist.HarpistDB.Save(&game).Error

	if err != nil {
		t.Error("Unable to add member to test ggameroup")
	}

	c := Character{Owner: TestUser, Name: "Test Character"}

	if harpist.HarpistDB.NewRecord(c) {
		harpist.HarpistDB.Create(&c)
		_ = harpist.HarpistDB.Save(&c).Error
	}

	game.AddCharacter(c)
	_ = harpist.HarpistDB.Save(&game).Error

	if !game.isPlayer(c) {
		t.Error("User did not store as a member of test game")
	}
}

func TestGameAddAdmin(t *testing.T) {
	game := Game{}

	harpist.HarpistDB.First(&game)
	game.AddAdmin(TestUser)
	err := harpist.HarpistDB.Save(&game).Error
	if err != nil {
		t.Error("Unable to add admin to test game")
	}
	var isAdmin bool

	for _, u := range game.GroupAdmins() {
		user := User{}
		harpist.HarpistDB.Where(u).First(&user)
		if u.LoginInfo.Username == user.LoginInfo.Username {
			isAdmin = true
		}
	}

	if !isAdmin {
		t.Error("User did not store as an admin of test game")
	}
}
