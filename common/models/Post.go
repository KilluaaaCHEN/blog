package models

type Post struct {
	PostId       int
	Title        string
	Tags         string
	Desc         string
	Content      string
	ReadCount    int
	CommentCount int
	TagCount     int
	IsOriginal   int

	TagsData []string
}

func (Post) TableName() string {
	return "post"
}
