package models

import (
	"time"
)

// Experience implements the in-memory representation of Experience
// TODO: Implement marshalling for storage
type Experience struct {
	changeAmount int
	logTime      time.Time
	next         *Experience
}

// ExperienceLog implements the in-memory representation of Experience logs
// TODO: Implement marshalling for storage
type ExperienceLog struct {
	ID        int
	character *Character
	length    int
	start     *Experience
}

// Append implemnt the ExperienceLog interface for appending new experience earnings
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

// Insert allows for inserting traits retroactively or as part of a rebuild of data
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

// TotalAtPoint allows for granular auditing of sheets against expenditures
// TODO: implement expenditures...
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

// Change implements convenience method for calculating the amount of XP earned over a period of time
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

// AuditedObject implements Audit interface
func (e ExperienceLog) AuditedObject() Character {
	return *e.character
}

// Attendance implements a relationship structure for game sessions
type Attendance struct {
	session *GameSession
	logTime time.Time
	next    *Attendance
}

// AttendanceLog implements log structure for character attendance to game sessions
type AttendanceLog struct {
	ID        int
	character *Character
	length    int
	start     *Attendance
}

// Append implements convenience method to append a new attendance entry to the log
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

// Insert implements convenience method to insert a new attendance record in its chronological index
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

// AuditedObject implements Audit interface
func (e AttendanceLog) AuditedObject() Harpist {
	return *e.character
}

// Approval implements generic approval type for later use in auditing
type Approval struct {
	ID       uint    `json:"id"`
	Owner    Harpist `json:"owner"`
	Approved bool    `json:"approved"`
}

// AuditedObject implements Audit interface
func (a Approval) AuditedObject() Harpist {
	return a.Owner
}

// IsApproved provides convenience method for checking an approval
func (a Approval) IsApproved() bool {
	return a.Approved
}

// Approve provides a convenience method for approving a requested approval
func (a *Approval) Approve() *Approval {
	if !a.IsApproved() {
		a.Approved = true
	}

	return a
}
