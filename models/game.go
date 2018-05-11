package models

import (
	"time"
)

// GameInfo implements base Game information used by interfaces
type GameInfo struct {
	ID      uint   `json:"id" yaml:"id" gorm:"primary_key"`
	Name    string `json:"name" yaml:"name"`
	Owner   User   `json:"owner" yaml:"owner" gorm:"foreignkey:OwnerID"`
	OwnerID uint   `json:"-" yaml:"-"`
}

// Game implements
type Game struct {
	ID         uint        `json:"id" yaml:"id" gorm:"primary_key"`
	Name       string      `json:"name" yaml:"name"`
	Owner      User        `json:"owner" yaml:"owner" gorm:"foreignkey:OwnerID"`
	OwnerID    uint        `json:"-" yaml:"-"`
	Characters []Character `json:"characters" yaml:"characters" gorm:"many2many:game_characters"`
	Players    []User      `json:"players" yaml:"players" gorm:"many2many:game_players"`
	GameAdmins []User      `json:"admins" yaml:"admins" gorm:"many2many:game_admins"`
}

// HarpistType implements...
func (g Game) HarpistType() HarpistType {
	return GAME
}

// Identity implements...
func (g Game) Identity() uint {
	return g.ID
}

func (g Game) isPlayer(c Character) bool {
	// TODO: goroutine this
	for _, ch := range g.Characters {
		if c == ch {
			return true
		}
	}
	return false
}

// GroupMembers implements...
func (g Game) GroupMembers() []Character {
	return g.Characters
}

// GroupAdmins implements...
func (g Game) GroupAdmins() []User {
	return g.GameAdmins
}

// GroupOwner implements...
func (g Game) GroupOwner() Harpist {
	return g.Owner
}

// MemberName implements...
func (g Game) MemberName() string {
	return g.Name
}

// MemberIdentity implements
func (g Game) MemberIdentity() uint {
	return g.ID
}

// AddAdmin implements
func (g *Game) AddAdmin(u User) error {
	for _, user := range g.GameAdmins {
		if user == u {
			return nil // Nothing to do, they're already there
		}
	}

	g.GameAdmins = append(g.GameAdmins, u)
	return nil
}

// RemoveAdmin implements
func (g *Game) RemoveAdmin(u User) error {
	for i, user := range g.GameAdmins {
		if user == u {
			g.GameAdmins = append(g.GameAdmins[:i], g.GameAdmins[i+1:]...)
			return nil
		}
	}
	return nil // nothing to do
}

// AddMember implements
func (g *Game) AddMember(u User) error {
	for _, user := range g.Players {
		if user == u {
			return nil // Nothing to do, they're already there
		}
	}

	g.Players = append(g.Players, u)
	return nil
}

// RemoveMember implements
func (g *Game) RemoveMember(u User) error {
	for i, user := range g.Players {
		if user == u {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
			return nil
		}
	}
	return nil // nothing to do
}

// AddCharacter implements
func (g *Game) AddCharacter(c Character) error {
	for _, char := range g.Characters {
		if char == c {
			return nil // Nothing to do, they're already there
		}
	}

	g.Characters = append(g.Characters, c)
	return nil
}

// RemoveCharacter implements
func (g *Game) RemoveCharacter(c Character) error {
	for i, char := range g.Characters {
		if char == c {
			g.Characters = append(g.Characters[:i], g.Characters[i+1:]...)
			return nil
		}
	}
	return nil // nothing to do
}

// Setting implements settings - groups of games that can have collections of characters
type Setting struct {
	ID         uint   `json:"id" yaml:"id" gorm:"primary_key"`
	Name       string `json:"name" yaml:"name"`
	Owner      User   `json:"owner" yaml:"owner" gorm:"foreignkey:OwnerID"`
	OwnerID    uint   `json:"-" yaml:"-"`
	Games      []Game `json:"games" yaml:"games"`
	GameAdmins []User `json:"setting_admins" yaml:"setting_admins"`
}

// HarpistType implements
func (s Setting) HarpistType() HarpistType {
	return SETTING
}

// Identity implements
func (s Setting) Identity() uint {
	return s.ID
}

func (s Setting) isPlayer(c Character) bool {
	for _, game := range s.GroupMembers() {
		if game.isPlayer(c) {
			return true
		}
	}
	return false
}

// GroupMembers implements
func (s Setting) GroupMembers() []Game {
	return s.Games
}

// GroupAdmins implements
func (s Setting) GroupAdmins() []User {
	return s.GameAdmins
}

// GroupOwner implements
func (s Setting) GroupOwner() Harpist {
	return s.Owner
}

// GameSession implements an event that can be used for auditing experience
type GameSession struct {
	HostingEntity *Group
	Game          *Group
	scheduledDate time.Time
	actualDate    time.Time
	canceled      bool
}

// GameObject implements
func (gs GameSession) GameObject() Group {
	return *gs.Game
}

// Info implements
func (gs GameSession) Info() Group {
	return *gs.Game
}

// Character implements basic chracter structure
// TODO: implement generic-ish sub-struct or map of optional sheet data for parsing by custom tools
type Character struct {
	ID      uint   `json:"id" yaml:"id" gorm:"primary_key"`
	Name    string `json:"name" yaml:"name"`
	Owner   User   `json:"owner" yaml:"owner" gorm:"foreignkey:OwnerID"`
	OwnerID uint   `json:"-" yaml:"-"`
}

// HarpistType implements
func (c Character) HarpistType() HarpistType {
	return CHARACTER
}

// Identity implements
func (c Character) Identity() uint {
	return c.ID
}

// MemberName implements
func (c Character) MemberName() string {
	return c.Owner.Name
}

// MemberIdentity implements
func (c Character) MemberIdentity() uint {
	return c.Owner.ID
}

// CharacterName implements convenience method for retrieving a character's name
func (c Character) CharacterName() string {
	return c.Name
}
