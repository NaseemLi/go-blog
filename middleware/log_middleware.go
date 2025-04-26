package middleware

import (
	logservice "goblog/service/log_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseWriter struct {
	gin.ResponseWriter
	Body []byte
	Head http.Header
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
	w.Body = append(w.Body, data...)
	return w.ResponseWriter.Write(data)
}

func (w *ResponseWriter) Header() http.Header {
	return w.Head
}

func LogMiddleware(c *gin.Context) {
	log := logservice.NewActionLogByGin(c)
	//请求中间件
	log.SetRequest(c)
	c.Set("log", log)
	if c.Request.URL.Path == "/api/ai/article" {
		c.Next()
		log.MiddlewareSave()
		return
	}

	res := &ResponseWriter{
		ResponseWriter: c.Writer,
		Head:           make(http.Header),
	}
	c.Writer = res
	c.Next()
	log.SetResponse(res.Body)
	log.SetResponseHeader(res.Head)
	log.MiddlewareSave()
}
