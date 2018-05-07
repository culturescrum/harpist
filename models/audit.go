package models

import (
	"time"
)

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

func (e *ExperienceLog) Insert(newEntry *Experience) {
	if e.length == 0 {
		e.start = newEntry
	} else {
		var previous *Experience
		current := e.start

		for !current.logTime.After(newEntry.logTime) {
			previous = current
			current = previous.next
		}

		previous.next = newEntry
		newEntry.next = current
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

type Attendance struct {
	session *GameSession
	logTime time.Time
	next    *Attendance
}

type AttendanceLog struct {
	ID        int
	character *Character
	length    int
	start     *Attendance
}

func (e *AttendanceLog) Append(newEntry *Attendance) {
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

func (e *AttendanceLog) Insert(newEntry *Attendance) {
	if e.length == 0 {
		e.start = newEntry
	} else {
		var previous *Attendance
		current := e.start

		for !current.logTime.After(newEntry.logTime) {
			previous = current
			current = previous.next
		}

		previous.next = newEntry
		newEntry.next = current
	}
	e.length++
}

func (a AttendanceLog) AuditedObject() Character {
	return *a.character
}

type Approval struct {
	ID    int64
	Owner Owner
}

func (a Approval) AuditedObject() Owner {
	return a.Owner
}
