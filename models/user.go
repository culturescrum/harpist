package models

// User defines basic user structure
type User struct {
  ID int64 `json:"id" yaml:"user_id"`
  Username string `json:"username" yaml:"username"`
  Password string `json:"password" yaml:"password"`
}
