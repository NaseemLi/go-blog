package res

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func SSEok(data any, c *gin.Context) {
	byteData, _ := json.Marshal(Response{SuccessCode, data, "成功"})
	c.SSEvent("", string(byteData))
	c.Writer.Flush()
}

func SSEfail(msg string, c *gin.Context) {
	byteData, _ := json.Marshal(Response{FailServiceCode, map[string]any{}, msg})
	c.SSEvent("", string(byteData))
	c.Writer.Flush()
}
