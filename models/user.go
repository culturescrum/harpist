package models

// User defines basic user structure
type User struct {
	ID   int64  `json:"id" yaml:"user_id"`
	Name string `json:"name" yaml:"name"`
}

type LoginInfo struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

func (u User) HarpistType() HarpistType {
	return USER
}

func (u User) Identity() int64 {
	return u.ID
}

func (u User) MemberName() string {
	return u.Name
}

func (u User) MemberIdentity() int64 {
	return u.ID
}
