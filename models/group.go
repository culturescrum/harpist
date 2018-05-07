package models

// PlayGroup defines the fields for generic groups
type PlayGroup struct {
	ID      uint   `json:"id" yaml:"id"`
	Name    string `json:"name" yaml:"name"`
	Owner   User   `json:"owner" yaml:"owner" gorm:"foreignkey:OwnerID"`
	OwnerID uint   `json:"-" yaml:"-"`
	Admins  []User `json:"-" yaml:"admins"`
	Members []User `json:"-" yaml:"members"`
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
