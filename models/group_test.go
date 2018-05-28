package models

import (
	"github.com/go-harpist/harpist"
	"testing"
)

var TestGroup PlayGroup

func init() {
	TestGroup = PlayGroup{
		Name: "Test Group",
	}

	harpist.HarpistDB.Where(TestUser).First(&TestUser)

	if harpist.HarpistDB.NewRecord(TestUser) {
		TestUser.SetPassword("password")
		harpist.HarpistDB.Create(&TestUser)
		_ = harpist.HarpistDB.Save(&TestUser).Error
	}

	if harpist.HarpistDB.NewRecord(TestGroup) {
		TestGroup.Owner = TestUser
		_ = harpist.HarpistDB.Save(&TestGroup).Error
	}

}

func TestGroupCreate(t *testing.T) {
	var err error

	harpist.HarpistDB.Where(TestGroup).First(&TestGroup)

	if harpist.HarpistDB.NewRecord(TestGroup) {
		TestGroup.Owner = TestUser
		err = harpist.HarpistDB.Save(&TestGroup).Error
	}

	if err != nil {
		t.Error("Unable to store group in the database")
	}

}

func TestGroupHarpistType(t *testing.T) {
	group := PlayGroup{}

	harpist.HarpistDB.Where(TestGroup).First(&group)
	if group.HarpistType() != GROUP {
		t.Error("HarpistType() does not return GROUP")
	}
}

func TestGroupIdentity(t *testing.T) {
	group := PlayGroup{}
	harpist.HarpistDB.First(&group)
	if group.Identity() != 1 {
		t.Error("Identity() does not test group ID 1")
	}
	if group.MemberIdentity() != 1 {
		t.Error("MemberIdentity() does not test group ID 1")
	}
}

func TestGroupOwner(t *testing.T) {
	group := PlayGroup{}
	harpist.HarpistDB.Where(TestGroup).First(&group)
	user := group.GroupOwner()
	harpist.HarpistDB.Where(user).First(&user)
	if group.GroupOwner().(User).ID != user.(User).ID {
		t.Error("Did not fetch correct user object")
	}
}

func TestGroupMemberName(t *testing.T) {
	group := PlayGroup{}

	harpist.HarpistDB.First(&group)
	if group.MemberName() != "Test Group" {
		t.Errorf("MemberName() does not return 'Test Group', returns %v", group.MemberName())
	}
}

func TestGroupAddMember(t *testing.T) {
	group := PlayGroup{}

	harpist.HarpistDB.Where(TestGroup).First(&group)
	group.AddMember(TestUser)
	err := harpist.HarpistDB.Save(&group).Error

	if err != nil {
		t.Error("Unable to add member to test group")
	}

	if !group.isPlayer(TestUser) {
		t.Error("User did not store as a member of test group")
	}
}

func TestGroupAddAdmin(t *testing.T) {
	group := PlayGroup{}

	harpist.HarpistDB.Where(TestGroup).First(&group)
	group.AddAdmin(TestUser)
	err := harpist.HarpistDB.Save(&group).Error
	if err != nil {
		t.Error("Unable to add admin to test group")
	}
	var isAdmin bool

	for _, u := range group.GroupAdmins() {
		user := User{}
		harpist.HarpistDB.Where(u).First(&user)
		if u.LoginInfo.Username == user.LoginInfo.Username {
			isAdmin = true
		}
	}

	if !isAdmin {
		t.Error("User did not store as an admin pf test group")
	}
}
