package middleware

import (
	"fmt"
	"goblog/global"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CacheOptions struct {
	Prefix CacheMiddlewarePrefix
	Time   time.Duration
	Params []string
}

type CacheMiddlewarePrefix string

const (
	CacheBannerPrefix CacheMiddlewarePrefix = "cache_banner_"
)

func NewBannerCacheOptions() CacheOptions {
	return CacheOptions{
		Prefix: CacheBannerPrefix,
		Time:   time.Hour * 60,
		Params: []string{"type"},
	}
}

type CacheResponseWriter struct {
	gin.ResponseWriter
	Body []byte
}

func (w *CacheResponseWriter) Write(data []byte) (int, error) {
	w.Body = append(w.Body, data...)
	return w.ResponseWriter.Write(data)
}

func CacheMiddleware(option CacheOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求
		values := url.Values{}
		for _, key := range option.Params {
			values.Add(key, c.Query(key))
		}
		key := fmt.Sprintf("%s%s", option.Prefix, values.Encode())
		fmt.Println(key)

		val, err := global.Redis.Get(key).Result()
		if err == nil {
			fmt.Printf("%s 使用缓存 %s\n", key, val)
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200, val)
			c.Abort()
			return
		}

		// 获取响应
		writer := &CacheResponseWriter{
			ResponseWriter: c.Writer,
		}
		c.Writer = writer
		c.Next()

		body := string(writer.Body)

		global.Redis.Set(key, body, option.Time)
	}
}

func CacheClose(prefix CacheMiddlewarePrefix) {
	keys, err := global.Redis.Keys(fmt.Sprintf("%s*", prefix)).Result()
	if err != nil {
		logrus.Errorf("Error retrieving keys: %v\n", err)
		return
	}

	if len(keys) > 0 {
		logrus.Infof("Deleting %d keys with prefix %s\n", len(keys), prefix)
		global.Redis.Del(keys...)
	}
}
