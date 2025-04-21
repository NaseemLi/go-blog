package focusservice

import (
	"goblog/global"
	"goblog/models"
	relationshipenum "goblog/models/enum/relationship_enum"
)

func CalcUserRelationship(A uint, B uint) (t relationshipenum.Relation) {
	var userFocusList []models.UserFocusModel
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

func CalcUserPatchRelationship(A uint, BList []uint) (m map[uint]relationshipenum.Relation) {
	m = make(map[uint]relationshipenum.Relation)
	var userFocusList []models.UserFocusModel
	global.DB.Find(&userFocusList, "(user_id = ? and focus_user_id in (?)) or (focus_user_id = ? and user_id in (?))", A, BList, A, BList)

	for _, B := range BList {
		//B 与 A 的关系
		var count int
		m[B] = relationshipenum.RelationStranger
		for _, model := range userFocusList {
			if model.FocusUserID == B {
				m[B] = relationshipenum.RelationFocus
				count++
			}
			if model.UserID == B {
				m[B] = relationshipenum.RelationFans
				count++
			}
		}
		if count == 2 {
			m[B] = relationshipenum.RelationFriend
		}
	}

	return
}
