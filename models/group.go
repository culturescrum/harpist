package models

// TemporalGroup defines the fields for generic groups
type TemporalGroup struct {
  ID            int64         `json:"id"`
  Name          string        `json:"name"`
  Creator       User          `json:"creator"`
  Admins        []User        `json:"-"`
  Members       []User        `json:"-"`
}

func (tg TemporalGroup) owner(user User) bool {
  return tg.Creator == user
}

func (tg TemporalGroup) isMember(user User) bool {
  for _, member := range tg.Members {
    if user == member {
      return true
    }
    // This will probably be controversial and up for discussion
    if user == tg.Creator {
      return true
    }
  }
  return false
}

// Conglomerate defines fields for an umbrella group owned by a group of
// users, mainly used to manage ConglomerateGroup structures.
type Conglomerate struct {
  TemporalGroup
  Directors     []User                `json:"-"`
  Members       []ConglomerateGroup   `json:"-"`
}

func (c Conglomerate) owner(user User) bool {
  for _, director := range c.Directors {
    if user == director {
      return true
    }
  }
  return false
}

func (c Conglomerate) isMember(user User) bool {
  for _, cg := range c.Members {
    if cg.isMember(user) {
      return true
    }
  }
  if c.owner(user) {
    return true
  }
  return false
}

// ConglomerateGroup defines fields for groups within an umbrella group
// with ownership passed to a Conglomerate
type ConglomerateGroup struct {
  TemporalGroup
  GroupName     string        `json:"group_name"`
  Owner         Conglomerate  `json:"-"`
}

func (cg ConglomerateGroup) owner(user User) bool {
  for _, director := range cg.Owner.Directors {
    if user == director {
      return true
    }
  }

  if cg.Creator == user {
    return true
  }

  return false
}
