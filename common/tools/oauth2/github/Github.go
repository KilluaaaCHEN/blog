package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Github struct {
	Id        int    `json:"id"`
	UserName  string `json:"login"`
	NickName  string `json:"name"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
	GitHubUrl string `json:"html_url"`
	BlogUrl   string `json:"blog"`
	Bio       string `json:"bio"`
	AppId     string
}

const (
	GetAtUrl       = "https://github.com/login/oauth/access_token?client_id=%v&client_secret=%v&code=%v"
	GetUserInfoUrl = "https://api.github.com/user?access_token=%v"
)

/**
获取Github用户信息
*/
func (u *Github) GetUserInfo(code string) (Github, error) {
	gh := Github{}
	ak, err := u.getAt(code)
	err = nil
	ak = "e9a5f559001ab702564d35ce9f53ea27e716825e"
	fmt.Println(ak)
	if err != nil {
		return gh, err
	}
	requestUrl := fmt.Sprintf(GetUserInfoUrl, ak)
	client := &http.Client{}
	resp, err := client.Get(requestUrl)
	if err != nil {
		return gh, err
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(result, &gh)
	if err != nil {
		return gh, err
	}
	gh.AppId = viper.GetString("oauth2.github.client_id")
	return gh, nil
}

/**
获取AccessToken
*/
func (u *Github) getAt(code string) (string, error) {

	githubConfig := viper.GetStringMapString("oauth2.github")

	requestUrl := fmt.Sprintf(GetAtUrl, githubConfig["client_id"], githubConfig["client_secret"], code)
	client := &http.Client{}
	resp, err := client.Get(requestUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	resultStr := string(result)
	resultMap, _ := url.ParseQuery(resultStr)
	if resultMap["error"] != nil {
		return "", errors.New(resultMap["error"][0] + "  " + resultMap["error_description"][0])
	}
	ak := resultMap["access_token"][0]
	return ak, nil
}
