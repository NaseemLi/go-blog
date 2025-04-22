package chatapi

import (
	"fmt"
	"goblog/common/res"
	"goblog/utils/jwts"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (ChatApi) ChatView(c *gin.Context) {
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil || claims == nil {
		res.FailWithMsg("token无效,请登录", c)
		return
	}
	conn, err := UP.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		//通常为客户端断开
		fmt.Println(err)
		return
	}

	for {
		// 消息类型，消息，错误
		t, p, err := conn.ReadMessage()
		if err != nil {
			if err == io.EOF {
				//客户端断开
				fmt.Println("客户端断开连接")
			}
			break
		}
		conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("你说的是：%s吗？", string(p))))
		fmt.Println(t, string(p))
	}
	defer conn.Close()
	fmt.Println("服务关闭")
}
