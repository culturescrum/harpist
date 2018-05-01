package models

// GameInfo implements base Game information used by interfaces
type GameInfo struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Owner Owner
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

func (g Game) Identity() int64 {
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

func (g Game) GroupOwner() Owner {
	return g.Owner
}

type Setting struct {
	GameInfo
	Games      []Game
	GameAdmins []User
}

func (s Setting) HarpistType() HarpistType {
	return SETTING
}

func (s Setting) Identity() int64 {
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

func (s Setting) GroupOwner() Owner {
	return s.Owner
}

type Character struct {
	GameInfo
}

func (c Character) HarpistType() HarpistType {
	return CHARACTER
}

func (c Character) Identity() int64 {
	return c.ID
}

func (c Character) MemberName() string {
	owner, _ := c.Owner.(User)
	return owner.Name
}

func (c Character) MemberIdentity() int64 {
	owner, _ := c.Owner.(User)
	return owner.ID
}

func (c Character) CharacterName() string {
	return c.Name
}
