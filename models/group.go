package models

// PlayGroup defines the fields for generic groups
type PlayGroup struct {
	ID      uint   `json:"id" yaml:"id" gorm:"primary_key"`
	Name    string `json:"name" yaml:"name"`
	Owner   User   `json:"owner" yaml:"owner" gorm:"foreignkey:OwnerID"`
	OwnerID uint   `json:"-" yaml:"-"`
	Admins  []User `json:"-" yaml:"admins" gorm:"many2many:group_admins"`
	Members []User `json:"-" yaml:"members" gorm:"many2many:group_members"`
}

func (PlayGroup) TableName() string {
	return "groups"
}

func (pg PlayGroup) HarpistType() HarpistType {
	return GROUP
}

func (pg PlayGroup) Identity() uint {
	return pg.ID
}

func (pg PlayGroup) isPlayer(u User) bool {
	for _, member := range pg.GroupMembers() {
		if u == member {
			return true
		}
	}
	return false
}

func (pg PlayGroup) GroupMembers() []User {
	return pg.Members
}

func (pg PlayGroup) GroupAdmins() []User {
	return pg.Admins
}

func (pg PlayGroup) GroupOwner() Harpist {
	return pg.Owner
}

func (pg PlayGroup) MemberName() string {
	return pg.Name
}

func (pg PlayGroup) MemberIdentity() uint {
	return pg.ID
}

func (pg *PlayGroup) AddAdmin(u User) error {
	for _, user := range pg.Admins {
		if user == u {
			return nil // Nothing to do, they're already there
		}
	}

	pg.Admins = append(pg.Admins, u)
	return nil
}

func (pg *PlayGroup) RemoveAdmin(u User) error {
	for i, user := range pg.Admins {
		if user == u {
			pg.Admins = append(pg.Admins[:i], pg.Admins[i+1:]...)
			return nil
		}
	}
	return nil // nothing to do
}

func (pg *PlayGroup) AddMember(u User) error {
	for _, user := range pg.Members {
		if user == u {
			return nil // Nothing to do, they're already there
		}
	}

	pg.Members = append(pg.Members, u)
	return nil
}

func (pg *PlayGroup) RemoveMember(u User) error {
	for i, user := range pg.Members {
		if user == u {
			pg.Members = append(pg.Members[:i], pg.Members[i+1:]...)
			return nil
		}
	}
	return nil // nothing to do
}

type Conglomerate struct {
	PlayGroup
	Members []PlayGroup `json:"-" yaml:"member_groups"`
}

func (c Conglomerate) HarpistType() HarpistType {
	return GROUP
}

func (c Conglomerate) Identity() uint {
	return c.ID
}

func (c Conglomerate) isPlayer(u User) bool {
	for _, subgroup := range c.GroupMembers() {
		for _, member := range subgroup.GroupMembers() {
			if u == member {
				return true
			}
		}
	}
	return false
}

func (c Conglomerate) GroupMembers() []PlayGroup {
	return c.Members
}

func (c Conglomerate) GroupOwner() Harpist {
	return c.Owner
}
