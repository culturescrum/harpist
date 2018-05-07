package models

type HarpistType int

const (
	USER HarpistType = iota
	GROUP
	CHARACTER
	GAME
	SETTING
)

type Owner interface {
	HarpistType() HarpistType
	Identity() int64
}

// Group creates the generic interface for groups and authorization
type Group interface {
	isPlayer(GroupMember) bool
	GroupMembers() []GroupMember
	GroupAdmins() []User
	GroupOwner() Owner
}

type GroupMember interface {
	MemberName() string
	MemberIdentity() int64
}

// TODO: Game models
// - [x] Game Data
// - [x] Character
// - [ ] Audit models (experience log, attendance, approvals)

// GameTier implements the interface for various collections in a game
type GameTier interface {
	GameObject() Group // Game / Setting
}

type CharacterType interface {
	CharacterName() string
}

type Audit interface {
	AuditedObject() Owner // returns Character for ExperienceLog, for example
}
