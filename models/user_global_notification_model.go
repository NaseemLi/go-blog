package models

type UserGlobalNotificationModel struct {
	Model
	NotificationID uint `json:"notificationID"` // 通知的唯一标识
	UserID         uint `json:"userID"`         // 用户 ID
	IsRead         bool `json:"isRead"`         // 是否已读
	IsDelete       bool `json:"isDelete"`       // 是否删除（逻辑删除）
}
