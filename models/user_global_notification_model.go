package models

type UserGlobalNotificationModel struct {
	Model
	NotificationID uint `json:"notification_id"` // 通知的唯一标识
	UserID         uint `json:"user_id"`         // 用户 ID
	IsRead         bool `json:"is_read"`         // 是否已读
	IsDelete       bool `json:"is_delete"`       // 是否删除（逻辑删除）
}
