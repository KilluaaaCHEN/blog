package models

import "gin/dao"

type Post struct {
	PostId    int `primary`
	Title     string
	Desc      string
	Content   string
	Tag       string
	ReadCount int
	TagCount  int
}

func (Post) TableName() string {
	return "post"
}

func (u *Post) GetPost(id int) Post {
	db := dao.Instance()
	var post Post
	db.First(&post, "post_id = ?", id)
	return post
}
