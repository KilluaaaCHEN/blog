package models

import "time"

type UserKey struct {
	Id        int
	Uid       int
	Type      string
	AppId     string
	Identity  string
	DeletedAt *time.Time
}

func (UserKey) TableName() string {
	return "user_key"
}
