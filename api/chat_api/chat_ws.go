package chatapi

import (
	"encoding/json"
	"fmt"
	"goblog/common/res"
	"goblog/global"
	"goblog/models"
	chatmsg "goblog/models/ctype/chat_msg"
	chatmsgtypeenum "goblog/models/enum/chat_msg_type_enum"
	"goblog/utils/jwts"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var OnlineMap = map[uint]map[string]*websocket.Conn{}

type ChatRequest struct {
	RevUserID uint                    `json:"revUserID"` //发给谁
	MsgType   chatmsgtypeenum.MsgType `json:"msgType"`   //1.文本 2.图片 3.Md
	Msg       chatmsg.ChatMsg         `json:"msg"`       //消息主体
}

type ChatResponse struct {
	ChatRecordResponse
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
		logrus.Errorf("WebSocket升级失败: %v", err)
		return
	}
	userID := claims.UserID
	var user models.UserModel
	err = global.DB.Take(&user, userID).Error
	if err != nil {
		res.SendConnFailWithMsg("用户不存在", conn)
		return
	}
	addr := conn.RemoteAddr().String()
	addrMap, ok := OnlineMap[userID]
	if !ok {
		OnlineMap[userID] = map[string]*websocket.Conn{
			addr: conn,
		}
	} else {
		_, ok1 := addrMap[addr]
		if !ok1 {
			OnlineMap[userID][addr] = conn
		}
	}
	//todo:未显示
	fmt.Sprintln("连接", OnlineMap)
	for {
		// 消息类型，消息，错误
		_, p, err := conn.ReadMessage()
		if err != nil {
			if err == io.EOF {
				//客户端断开
				fmt.Println("客户端断开连接")
			}
			break
		}

		var req ChatRequest
		err2 := json.Unmarshal(p, &req)
		if err2 != nil {
			res.SendConnFailWithMsg("消息格式错误", conn)
			continue
		}
		//判断接收人在不在
		var revUserID models.UserModel
		err = global.DB.Take(&revUserID, req.RevUserID).Error
		if err != nil {
			res.SendConnFailWithMsg("接收人不存在", conn)
			continue
		}
		//先落库
		model := models.ChatModel{
			SendUserID: claims.UserID,
			RevUserID:  req.RevUserID,
			MsgType:    req.MsgType,
			Msg:        req.Msg,
		}
		err = global.DB.Create(&model).Error
		if err != nil {
			res.SendConnFailWithMsg("数据库保存失败", conn)
			continue
		}

		item := ChatResponse{
			ChatRecordResponse: ChatRecordResponse{
				ChatModel:        model,
				SendUserNickname: user.Nickname,
				SendUserAvatar:   user.Avatar,
				RevUserNickname:  revUserID.Nickname,
				RevUserAvatar:    revUserID.Avatar,
			},
		}
		//消息接收人,看看是否在线,在线就发送给对方
		//发给对方
		res.SendWsMsg(OnlineMap, req.RevUserID, item)
		//发给自己
		item.IsMe = true
		res.SendConnOkWithData(item, conn)
	}

	defer conn.Close()
	// 用户离开的时候
	addrMap2, ok2 := OnlineMap[userID]
	if ok2 {
		_, ok3 := addrMap2[addr]
		if ok3 {
			delete(OnlineMap[userID], addr)
		}
		if len(addrMap2) == 0 {
			delete(OnlineMap, userID)
		}
	}
	fmt.Sprintln("离开", OnlineMap)
}
