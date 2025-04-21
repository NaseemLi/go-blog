package models

import (
	chatmsg "goblog/models/ctype/chat_msg"
	chatmsgtypeenum "goblog/models/enum/chat_msg_type_enum"
)

type ChatModel struct {
	Model
	SendUserID    uint                    `json:"sendUserID"`                               // 发送用户ID
	SendUserModel UserModel               `gorm:"foreignKey:SendUserID" json:"-"`           // 发送用户
	RevUserID     uint                    `json:"revUserID"`                                // 接收用户ID
	RevUserModel  UserModel               `gorm:"foreignKey:RevUserID" json:"-"`            // 接收用户
	MsgType       chatmsgtypeenum.MsgType `json:"msgType"`                                  // 消息类型
	Msg           chatmsg.ChatMsg         `gorm:"type:longtext;serializer:json" json:"msg"` // 消息内容
}
