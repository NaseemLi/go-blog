package middleware

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogMiddleware(c *gin.Context) {
	//请求中间件
	byteData, err := io.ReadAll(c.Request.Body) //阅后即焚
	if err != nil {
		logrus.Errorf(err.Error())
	}
	fmt.Println("body: ", string(byteData))
	c.Request.Body = io.NopCloser(bytes.NewReader(byteData))
	c.Next()
	//相应中间件

}
