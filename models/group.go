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

// TableName provides the name for gorm
func (PlayGroup) TableName() string {
	return "groups"
}

// HarpistType implements Harpist interface
func (pg PlayGroup) HarpistType() HarpistType {
	return GROUP
}

// Identity implement Harpist interface
func (pg PlayGroup) Identity() uint {
	return pg.ID
}

// isPlayer implements Group interface
func (pg PlayGroup) isPlayer(u User) bool {
	for _, member := range pg.GroupMembers() {
		if u == member {
			return true
		}
	}
	return false
}

// GroupMembers imolements Group interface
func (pg PlayGroup) GroupMembers() []User {
	return pg.Members
}

// GroupAdmins implements Group interface
func (pg PlayGroup) GroupAdmins() []User {
	return pg.Admins
}

// GroupOwner implements Group interface
func (pg PlayGroup) GroupOwner() Harpist {
	return pg.Owner
}

// MemberName implements GroupMember interface
func (pg PlayGroup) MemberName() string {
	return pg.Name
}

// MemberIdentity implements GroupMember interface
func (pg PlayGroup) MemberIdentity() uint {
	return pg.ID
}

// AddAdmin implements convenience method for adding admins to a Group
// TODO: implement anonymous function that can handle conglomerates
func (pg *PlayGroup) AddAdmin(u User) error {
	for _, user := range pg.Admins {
		if user == u {
			return nil // Nothing to do, they're already there
		}
	}

	pg.Admins = append(pg.Admins, u)
	return nil
}

// RemoveAdmin implements convenience method for removing admins from a Group
// TODO: implement anonymous function that can handle conglomerates
func (pg *PlayGroup) RemoveAdmin(u User) error {
	for i, user := range pg.Admins {
		if user == u {
			pg.Admins = append(pg.Admins[:i], pg.Admins[i+1:]...)
			return nil
		}
	}
	return nil // nothing to do
}

// AddMember implments convenience method for adding members to a Group
// TODO: implement anonymous function that can handle conglomerate sub-group memberships
func (pg *PlayGroup) AddMember(u User) error {
	for _, user := range pg.Members {
		if user == u {
			return nil // Nothing to do, they're already there
		}
	}

	pg.Members = append(pg.Members, u)
	return nil
}

// RemoveMember implements convenience method for removing memobers from a Group
// TODO: implement anonymous fnction that can handle removing sub-groups from conglomerates
func (pg *PlayGroup) RemoveMember(u User) error {
	for i, user := range pg.Members {
		if user == u {
			pg.Members = append(pg.Members[:i], pg.Members[i+1:]...)
			return nil
		}
	}
	return nil // nothing to do
}

// Conglomerate implements gaming organizations that have sub-groups
// TODO: implement dynamic grouping trees (sigh)
type Conglomerate struct {
	PlayGroup
	Members []PlayGroup `json:"-" yaml:"member_groups"`
}

// HarpistType implements Harpist interface
func (c Conglomerate) HarpistType() HarpistType {
	return GROUP
}

// Identity implements Harpist interface
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

// GroupMembers implements Group interface
func (c Conglomerate) GroupMembers() []PlayGroup {
	return c.Members
}

// GroupOwner implements Group interface
func (c Conglomerate) GroupOwner() Harpist {
	return c.Owner
}

// TODO: implement the rest of Group interface for Conglomerate
