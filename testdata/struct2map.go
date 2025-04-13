package main

import (
	"fmt"
)

type UserInfoUpateRequest struct {
	Username    *string   `json:"username" s-u:"username"`           // 用户名
	Nickname    *string   `json:"nickname" s-u:"nickname"`           // 昵称
	Avatar      *string   `json:"avatar" s-u:"avatar"`               // 头像
	Abstract    *string   `json:"abstract" s-u:"abstract"`           // 简介
	LikeTags    *[]string `json:"likeTags" s-u-c:"like_tags"`        // 喜欢的标签
	OpenCollect *bool     `json:"openCollect" s-u-c:"open_collect"`  // 是否开启收藏
	OpenFollow  *bool     `json:"openFollow" s-u-c:"open_follow"`    // 是否开启关注
	OpenFans    *bool     `json:"openFans" s-u-c:"open_fans"`        // 是否开启粉丝
	HomeStyleID *uint     `json:"homeStyleID" s-u-c:"home_style_id"` // 首页推荐样式id
}

func main() {
	var name = "那没事儿"
	var openFans = true
	var cr = UserInfoUpateRequest{
		Nickname: &name,
		OpenFans: &openFans,
	}
	fmt.Println(Struct2map(cr, "s-u"))
	fmt.Println(Struct2map(cr, "s-u-c"))
}
