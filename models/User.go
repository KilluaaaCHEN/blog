package models

import (
	"gin/dao"
	_ "github.com/jinzhu/gorm"
)

type User struct {
	Id       int
	Nickname string `json:"nick_name"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) GetUser(id int) User {
	db := dao.Instance()
	var user User
	db.First(&user, id)
	return user
}
