package models

// Group creates the generic interface for groups and authorization
type Group interface {
  owner(User)               bool
  isMember(User)            bool
}

// TODO: Game models
// - [ ] Game Data
// - [ ] Character
// - [ ] Audit models (experience log, attendance, approvals)
// - [ ] External Data (OAuth, URLs)
// - [ ] Events

// GameTier implements the interface for various collections in a game
type GameTier interface {
}
