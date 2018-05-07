package models

import (
	"time"
)

// GameInfo implements base Game information used by interfaces
type GameInfo struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Owner User
}

// Game implements
type Game struct {
	GameInfo
	Characters []Character
	Visitors   []Character
	GameAdmins []User
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
	for _, ch := range g.Visitors {
		if c == ch {
			return true
		}
	}
	return false
}

func (g Game) GroupMembers() []Character {
	return append(g.Characters, g.Visitors...)
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

type Setting struct {
	GameInfo
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
	GameInfo
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
