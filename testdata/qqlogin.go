package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"goblog/global"
	"io"
	"net/http"
	"net/url"
)

type AccessTokenType struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	Openid       string `json:"openid"`
	RefreshToken string `json:"refresh_token"`
}

func GetAccessToken(code string) (at AccessTokenType, err error) {
	qq := global.Config.QQ
	baseUrl := url.Parse("https://graph.qq.com/oauth2.0/token")

	p := url.Values{}
	p.Add("grant_type", "authorization_code")
	p.Add("client_id", qq.AppID)
	p.Add("client_secret", qq.AppKey)
	p.Add("code", code)
	p.Add("fmt", "json")
	p.Add("redirect_uri", qq.Redirect)
	p.Add("need_openid", "1")
	baseUrl.RawQuery = p.Encode()
	resposne, err := http.Get(baseUrl.String())
	if err != nil {
		return
	}
	byteData, _ := io.ReadAll(resposne.Body)
	err = json.Unmarshal(byteData, &at)
	if err != nil {
		return
	}
	if at.AccessToken == "" {
		fmt.Sprintf(string(byteData))
		err = errors.New("获取 accesstoken 失败")
		return
	}
	return
}

type UserInfo struct {
	Ret             int    `json:"ret"`
	Msg             string `json:"msg"`
	IsLost          int    `json:"is_lost"`
	Nickname        string `json:"nickname"`
	Gender          string `json:"gender"`
	GenderType      string `json:"gender_type"`
	Province        string `json:"province"`
	City            string `json:"city"`
	Year            string `json:"year"`
	Figureurl       string `json:"figureurl"`
	Figureurl1      string `json:"figureurl_1"`
	Figureurl2      string `json:"figureurl_2"`
	FigureurlQQ1    string `json:"figureurl_qq_1"`
	FigureurlQQ2    string `json:"figureurl_qq_2"`
	IsYellowVIP     string `json:"is_yellow_vip"`
	VIP             string `json:"vip"`
	YellowVIPLevel  string `json:"yellow_vip_level"`
	Level           string `json:"level"`
	IsYellowYearVIP string `json:"is_yellow_year_vip"`
}

func GetUserInfo(at AccessTokenType) (userinfo UserInfo, err error) {
	qq := global.Config.QQ
	baseUrl := url.Parse()

	p := url.Values{}
	p.Add("access_token", at.AccessToken)
	p.Add("oauth_consumer_key", qq.AppID)
	p.Add("openid", at.Openid)

	baseUrl.RawQuery = p.Encode()
	resposne, err := http.Get(baseUrl.String())
	if err != nil {
		return
	}
	byteData, _ := io.ReadAll(resposne.Body)
	err = json.Unmarshal(byteData, &userinfo)
	if err != nil {
		return
	}
	if userinfo.Ret != 0 {
		err = errors.New(userinfo.Msg)
		return
	}
	return
}

func main() {
}
