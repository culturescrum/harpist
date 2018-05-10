package models

import (
	"time"
)

// GameInfo implements base Game information used by interfaces
type GameInfo struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Owner   User   `json:"owner" gorm:"foreignkey:OwnerID"`
	OwnerID uint   `json:"-"`
}

// Game implements
type Game struct {
	ID         uint        `json:"id" gorm:"primary_key"`
	Name       string      `json:"name"`
	Owner      User        `json:"owner" gorm:"foreignkey:OwnerID"`
	OwnerID    uint        `json:"-"`
	Characters []Character `json:"characters" gorm:"many2many:game_characters"`
	Players    []User      `json:"players" gorm:"many2many:game_players"`
	GameAdmins []User      `json:"admins" gorm:"many2many:game_admins"`
}

func (g Game) HarpistType() HarpistType {
	return GAME
}

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

func (g Game) GroupMembers() []Character {
	return g.Characters
}

func (g Game) GroupAdmins() []User {
	return g.GameAdmins
}

func (g Game) GroupOwner() Harpist {
	return g.Owner
}

func (g Game) MemberName() string {
	return g.Name
}

func (g Game) MemberIdentity() uint {
	return g.ID
}

func (g *Game) AddAdmin(u User) error {
	for _, user := range g.GameAdmins {
		if user == u {
			return nil // Nothing to do, they're already there
		}
	}

	g.GameAdmins = append(g.GameAdmins, u)
	return nil
}

func (g *Game) RemoveAdmin(u User) error {
	for i, user := range g.GameAdmins {
		if user == u {
			g.GameAdmins = append(g.GameAdmins[:i], g.GameAdmins[i+1:]...)
			return nil
		}
	}
	return nil // nothing to do
}

func (g *Game) AddMember(u User) error {
	for _, user := range g.Players {
		if user == u {
			return nil // Nothing to do, they're already there
		}
	}

	g.Players = append(g.Players, u)
	return nil
}

func (g *Game) RemoveMember(u User) error {
	for i, user := range g.Players {
		if user == u {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
			return nil
		}
	}
	return nil // nothing to do
}

func (g *Game) AddCharacter(c Character) error {
	for _, char := range g.Characters {
		if char == c {
			return nil // Nothing to do, they're already there
		}
	}

	g.Characters = append(g.Characters, c)
	return nil
}

func (g *Game) RemoveCharacter(c Character) error {
	for i, char := range g.Characters {
		if char == c {
			g.Characters = append(g.Characters[:i], g.Characters[i+1:]...)
			return nil
		}
	}
	return nil // nothing to do
}

type Setting struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Name       string `json:"name"`
	Owner      User   `json:"owner" gorm:"foreignkey:OwnerID"`
	OwnerID    uint   `json:"-"`
	Games      []Game
	GameAdmins []User
}

func (s Setting) HarpistType() HarpistType {
	return SETTING
}

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

func (s Setting) GroupMembers() []Game {
	return s.Games
}

func (s Setting) GroupAdmins() []User {
	return s.GameAdmins
}

func (s Setting) GroupOwner() Harpist {
	return s.Owner
}

type GameSession struct {
	HostingEntity *Group
	Game          *Group
	scheduledDate time.Time
	actualDate    time.Time
	canceled      bool
}

func (gs GameSession) GameObject() Group {
	return *gs.Game
}

func (gs GameSession) Info() Group {
	return *gs.Game
}

type Character struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Owner   User   `json:"owner" gorm:"foreignkey:OwnerID"`
	OwnerID uint   `json:"-"`
}

func (c Character) HarpistType() HarpistType {
	return CHARACTER
}

func (c Character) Identity() uint {
	return c.ID
}

func (c Character) MemberName() string {
	return c.Owner.Name
}

func (c Character) MemberIdentity() uint {
	return c.Owner.ID
}

func (c Character) CharacterName() string {
	return c.Name
}
