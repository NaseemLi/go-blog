package chatapi

import (
	"encoding/json"
	"fmt"
	"goblog/common/res"
	"goblog/global"
	"goblog/models"
	chatmsg "goblog/models/ctype/chat_msg"
	chatmsgtypeenum "goblog/models/enum/chat_msg_type_enum"
	relationshipenum "goblog/models/enum/relationship_enum"
	focusservice "goblog/service/focus_service"
	"goblog/utils/jwts"
	"goblog/utils/xss"
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

		// 具体的消息类型做处理
		switch req.MsgType {

		case chatmsgtypeenum.TextMsgType:
			if req.Msg.TextMsg == nil || req.Msg.TextMsg.Content == "" {
				res.SendConnFailWithMsg("文本消息内容为空", conn)
				continue
			}
		case chatmsgtypeenum.ImageMsgType:
			if req.Msg.ImageMsg == nil || req.Msg.ImageMsg.Src == "" {
				res.SendConnFailWithMsg("图片消息内容为空", conn)
				continue
			}
		case chatmsgtypeenum.MarkDownMsgType:
			if req.Msg.MarkDownMsg == nil || req.Msg.MarkDownMsg.Content == "" {
				res.SendConnFailWithMsg("markdown消息内容为空", conn)
				continue
			}
			// 对markdown消息做过滤
			req.Msg.MarkDownMsg.Content = xss.XSSFilter(req.Msg.MarkDownMsg.Content)
		default:
			res.SendConnFailWithMsg("不支持的消息类型", conn)
			continue
		}

		// 判断你与对方的好友关系
		// 好友就能每天聊
		// 已关注和粉丝 如果对方没有回复你，那么每天只能聊一次  对方如果没有回你，那么你只能发三条消息
		// 陌生人，如果对方开了陌生人私信，那么就能聊
		relation := focusservice.CalcUserRelationship(userID, req.RevUserID)

		switch relation {
		case relationshipenum.RelationStranger: // 陌生人
			var revUserMsgConf models.UserMessageConfModel
			err = global.DB.Take(&revUserMsgConf, "user_id = ?", revUserID.ID).Error
			if err != nil {
				res.SendConnFailWithMsg("接收人隐私设置不存在", conn)
				continue
			}
			if !revUserMsgConf.OpenPrivateChat {
				res.SendConnFailWithMsg("对方未开始陌生人私聊", conn)
				continue
			}
			// 陌生人频控：对方未回复前，每天只能发一条
			var chatList []models.ChatModel
			global.DB.Find(&chatList, "date(created_at) = date (now()) and ( (send_user_id = ? and  rev_user_id = ?) or (send_user_id = ? and  rev_user_id = ?))",
				userID, req.RevUserID, req.RevUserID, userID)

			var sendChatCount, revChatCount int
			for _, model := range chatList {
				if model.SendUserID == userID {
					sendChatCount++
				}
				if model.RevUserID == userID {
					revChatCount++
				}
			}
			if sendChatCount > 0 && revChatCount == 0 {
				res.SendConnFailWithMsg("对方未回复的情况下，陌生人每天只能发送一条消息", conn)
				continue
			}
		case relationshipenum.RelationFocus, relationshipenum.RelationFans: // 已关注
			// 今天对方如果没有回复你，那么你就只能发一条
			var chatList []models.ChatModel
			global.DB.Find(&chatList, "date(created_at) = date (now()) and ( (send_user_id = ? and  rev_user_id = ?) or (send_user_id = ? and  rev_user_id = ?))",
				userID, req.RevUserID, req.RevUserID, userID)

			// 我发的  对方发的
			var sendChatCount, revChatCount int
			for _, model := range chatList {
				if model.SendUserID == userID {
					sendChatCount++
				}
				if model.RevUserID == userID {
					revChatCount++
				}
			}

			if sendChatCount > 0 && revChatCount == 0 {
				res.SendConnFailWithMsg("对方未回复的情况下，当天只能发送一条消息", conn)
				continue
			}

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
