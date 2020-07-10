package models

import (
	"blog_api/common/tools"
	"blog_api/common/tools/oauth2/github"
	"blog_api/dao"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type User struct {
	Id        int
	GithubId  string
	UserName  string
	Password  string
	NickName  string `gorm:"column:nickname"`
	Email     string
	AvatarUrl string
	GitHubUrl string `gorm:"column:github_url"`
	BlogUrl   string
	Bio       string
	State     int8
	LastAt    int
	LastIp    string
	RegIp     string
	CreatedAt int
	UpdatedAt int
	DeletedAt *time.Time
}

func (User) TableName() string {
	return "users"
}

const (
	StateActive = 1
)

/**
用户登陆
*/
func (u User) Login(c *gin.Context, ip string) error {
	if u.State != StateActive {
		return errors.New("账号违规,已被停用")
	}
	db := dao.InstDB()
	db.Model(&u).Updates(User{LastAt: int(time.Now().Unix()), LastIp: ip})
	return nil
}

/**
Github登录
*/
func (u User) LoginGithubUser(c *gin.Context, gh github.Github, ip string) (User, error) {
	user := User{}
	if gh.Id == 0 {
		return user, errors.New("获取账号失败")
	}
	db := dao.InstDB()
	if db.Where("github_id = ?", gh.Id).First(&user).RecordNotFound() {
		user = User{
			GithubId:  strconv.Itoa(gh.Id),
			UserName:  gh.UserName,
			NickName:  gh.NickName,
			Email:     gh.Email,
			AvatarUrl: gh.AvatarUrl,
			GitHubUrl: gh.GitHubUrl,
			BlogUrl:   gh.BlogUrl,
			Bio:       gh.Bio,
			State:     1,
			RegIp:     ip,
			Password:  tools.Str{}.Random(128),
		}
		db.Create(&user)
	}
	if user.Id == 0 {
		return user, errors.New("账号创建失败")
	}
	err := user.Login(c, ip)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *User) GetUserByPwd(pwd string) (User, error) {
	db := dao.InstDB()
	user := User{}
	db.Where("password = ?", pwd).First(&user)
	if user.Id == 0 {
		return user, errors.New("请重新登录")
	}
	return user, nil
}
