package models

// PlayGroup defines the fields for generic groups
type PlayGroup struct {
	ID      int64  `json:"id" yaml:"id"`
	Name    string `json:"name" yaml:"name"`
	Owner   Owner  `json:"owner" yaml:"owner"`
	Admins  []User `json:"-" yaml:"admins"`
	Members []User `json:"-" yaml:"members"`
}

func (pg PlayGroup) HarpistType() HarpistType {
	return GROUP
}

func (pg PlayGroup) Identity() int64 {
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

func (pg PlayGroup) GroupOwner() Owner {
	return pg.Owner
}

func (pg PlayGroup) MemberName() string {
	return pg.Name
}

func (pg PlayGroup) MemberIdentity() int64 {
	return pg.ID
}

type Conglomerate struct {
	PlayGroup
	Members []PlayGroup `json:"-" yaml:"member_groups"`
}

func (c Conglomerate) HarpistType() HarpistType {
	return GROUP
}

func (c Conglomerate) Identity() int64 {
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

func (c Conglomerate) GroupOwner() Owner {
	return c.Owner
}
