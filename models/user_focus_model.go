package models

type UserFocusModel struct {
	Model
	UserID         uint      `json:"userID"` //用户ID
	UserModel      UserModel `gorm:"foreignKey:UserID" json:"-"`
	FocusUserID    uint      `json:"focusUserID"` //关注的用户
	FocusUserModel UserModel `gorm:"foreignKey:FocusUserID" json:"-"`
}
