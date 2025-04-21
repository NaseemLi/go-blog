package models

import (
	"goblog/global"
	relationshipenum "goblog/models/enum/relationship_enum"
)

type UserFocusModel struct {
	Model
	UserID         uint      `json:"userID"` //用户ID
	UserModel      UserModel `gorm:"foreignKey:UserID" json:"-"`
	FocusUserID    uint      `json:"focusUserID"` //关注的用户
	FocusUserModel UserModel `gorm:"foreignKey:FocusUserID" json:"-"`
}

func CalcUserRelationship(A uint, B uint) (t relationshipenum.Relation) {
	var userFocusList []UserFocusModel
	global.DB.Find(&userFocusList, "(user_id = ? and focus_user_id = ?) or (focus_user_id = ? and user_id = ?)", A, B, B, A)
	if len(userFocusList) == 2 {
		return relationshipenum.RelationFriend
	}
	if len(userFocusList) == 0 {
		return relationshipenum.RelationStranger
	}
	focus := userFocusList[0]
	if focus.UserID == A {
		return relationshipenum.RelationFocus
	}

	return relationshipenum.RelationFans
}
