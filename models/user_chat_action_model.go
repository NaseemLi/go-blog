package models

type UserChatActionModel struct {
	Model
	UserID   uint `json:"userID"` // 用户ID
	ChatID   uint `json:"chatID"` // 聊天记录ID
	IsRead   bool `json:"isRead"`
	IsRelete bool `json:"isRelete"`
}
