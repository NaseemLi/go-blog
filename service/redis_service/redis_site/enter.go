package redissite

import "goblog/global"

const key = "blog_site_flow"

func SetFlow() {
	v, _ := global.Redis.Get(key).Int()
	global.Redis.Set("blog_site_flow", v+1, 0)
}

func GetFlow() int {
	v, _ := global.Redis.Get(key).Int()
	return v
}
