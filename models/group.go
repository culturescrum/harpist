package models

// Group contains the fields for inheritance in other group types
type Group struct {
  ID            int64         `json:"id"`
  Name          string        `json:"name"`
}

// TemporalGroup defines fields for a centralized group owned by a
// single user
type TemporalGroup struct {
  Group
  Owner         User          `json:"owner"`
  Members       []User        `json:"-"`
}

// Conglomerate defines fields for an umbrella group owned by a group of
// users, mainly used to manage ConglomerateGroup structures.
type Conglomerate struct {
  Group
  Owners        []User        `json:"-"`
}

// ConglomerateGroup defines fields for groups within an umbrella group
// with ownership passed to a Conglomerate
type ConglomerateGroup struct {
  Group
  Owner         Conglomerate  `json:"-"`
  Admins        []User        `json:"-"`
}
