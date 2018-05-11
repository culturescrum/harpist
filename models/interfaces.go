package models

// HarpistType const convenience
type HarpistType int

const (
	// USER is a user
	USER HarpistType = iota
	// GROUP is a group or conglomerate (TODO: interfaces)
	GROUP
	// CHARACTER is a character of some kind
	CHARACTER
	// GAME is a troupe or gaming group's actual game
	GAME
	// SETTING is a collection of troupes or gaming groups' games
	SETTING
)

// Harpist interface to be used for anonymizing many functions at a later date
type Harpist interface {
	HarpistType() HarpistType
	Identity() uint
}

// Group creates the generic interface for groups and authorization
type Group interface {
	isPlayer(GroupMember) bool
	GroupMembers() []GroupMember
	GroupAdmins() []User
	GroupOwner() Harpist
}

// GroupMember interface to be used for anonymizing membership functions at a later date
type GroupMember interface {
	MemberName() string
	MemberIdentity() uint
}

// GameTier implements the interface for various collections in a game
type GameTier interface {
	GameObject() Group // Game / Setting
}

// CharacterType is a starting point for a full interface for character sheets
type CharacterType interface {
	CharacterName() string
}

// Audit types are a thing
type Audit interface {
	AuditedObject() Harpist // returns Character for ExperienceLog, for example
}
