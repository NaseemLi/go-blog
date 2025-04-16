package redisarticle

import (
	"goblog/global"
	"strconv"
)

type articleCacheType string

const (
	articleCacheLook    articleCacheType = "article_look_key"
	articleCacheDigg    articleCacheType = "article_digg_key"
	articleCacheCollect articleCacheType = "article_collect_key"
)

func set(t articleCacheType, articleID uint, increase bool) {
	num, _ := global.Redis.HGet(string(t), strconv.Itoa(int(articleID))).Int()
	if increase {
		num--
	} else {
		num++
	}
	global.Redis.HSet(string(t), strconv.Itoa(int(articleID)), num)
}

func SetCacheLook(articleID uint, increase bool) {
	set(articleCacheLook, articleID, increase)
}

func SetCacheDigg(articleID uint, increase bool) {
	set(articleCacheDigg, articleID, increase)
}

func SetCacheCollect(articleID uint, increase bool) {
	set(articleCacheCollect, articleID, increase)
}
