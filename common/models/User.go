package models

import (
	"blog/common/tools/oauth2/github"
	"blog/dao"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type User struct {
	Id        int
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
	u.StoreLogin(c)
	return nil
}

/**
Github登录
*/
func (u User) LoginGithubUser(c *gin.Context, gh github.Github, ip string) (User, error) {
	db := dao.InstDB()
	var userKey UserKey
	user := User{}
	if db.Where("type = 'github' AND app_id = ? AND identity = ?", gh.AppId, gh.Id).First(&userKey).RecordNotFound() {
		user := User{
			UserName:  gh.UserName,
			NickName:  gh.NickName,
			Email:     gh.Email,
			AvatarUrl: gh.AvatarUrl,
			GitHubUrl: gh.GitHubUrl,
			BlogUrl:   gh.BlogUrl,
			Bio:       gh.Bio,
			State:     1,
			RegIp:     ip,
		}
		db.Create(&user)

		userKey.AppId = gh.AppId
		userKey.Type = "github"
		userKey.Identity = strconv.Itoa(gh.Id)
		userKey.Uid = user.Id
		db.Create(&userKey)
	} else {
		db.First(&user, userKey.Uid)
	}
	err := user.Login(c, ip)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

/**
保存登录凭证
*/
func (u *User) StoreLogin(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("user_id", u.Id)
	session.Set("user_name", u.UserName)
	session.Save()
}
