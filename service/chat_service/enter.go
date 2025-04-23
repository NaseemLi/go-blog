package chatservice

import (
	"goblog/global"
	"goblog/models"
	chatmsg "goblog/models/ctype/chat_msg"
	chatmsgtypeenum "goblog/models/enum/chat_msg_type_enum"
	"goblog/utils/xss"

	"github.com/sirupsen/logrus"
)

// A -> B发消息
func ToChat(A, B uint, msgType chatmsgtypeenum.MsgType, msg chatmsg.ChatMsg) {
	err := global.DB.Create(&models.ChatModel{
		SendUserID: A,
		RevUserID:  B,
		MsgType:    msgType,
		Msg:        msg,
	}).Error
	if err != nil {
		logrus.Errorf("对话创建失败: %v", err)
		return
	}
}

func ToTextChat(A, B uint, content string) {
	ToChat(A, B, chatmsgtypeenum.TextMsgType, chatmsg.ChatMsg{
		TextMsg: &chatmsg.TextMsg{
			Content: content,
		},
	})
}

func ToImageChat(A, B uint, src string) {
	ToChat(A, B, chatmsgtypeenum.ImageMsgType, chatmsg.ChatMsg{
		ImageMsg: &chatmsg.ImageMsg{
			Src: src,
		},
	})
}

func ToMarkDownChat(A, B uint, content string) {
	//过滤xss
	filterContent := xss.XSSFilter(content)

	ToChat(A, B, chatmsgtypeenum.MarkDownMsgType, chatmsg.ChatMsg{
		MarkDownMsg: &chatmsg.MarkDownMsg{
			Content: filterContent,
		},
	})
}
