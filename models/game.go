package models

import (
	"time"
)

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

type Experience struct {
	changeAmount int
	logTime      time.Time
	next         *Experience
}

type ExperienceLog struct {
	ID        int
	character *Character
	length    int
	start     *Experience
}

func (e *ExperienceLog) Append(newEntry *Experience) {
	if e.length == 0 {
		e.start = newEntry
	} else {
		current := e.start
		for current.next != nil {
			current = current.next
		}
		current.next = newEntry
	}
	e.length++
}

func (e ExperienceLog) TotalAtPoint(t time.Time) int {
	var (
		total int
	)
	current := e.start
	// As long as the log time isn't after the given time...
	// TODO: make this more forgiving than microseconds.
	for current.next != nil {
		if !current.logTime.After(t) {
			total += current.changeAmount
		} else {
			break
		}
		current = current.next
	}
	return total
}

func (e ExperienceLog) Change(t time.Time, u time.Time) int {
	var (
		total int
	)
	current := e.start
	for current.next != nil {
		if u.After(current.logTime) {
			if current.logTime.After(t) {
				total += current.changeAmount
				current = current.next
			} else {
				current = current.next
				continue
			}
		} else {
			break
		}
	}
	return total
}

func (e ExperienceLog) AuditedObject() Character {
	return *e.character
}
