package models

type UserMessageConfModel struct {
	UserID             uint      `gorm:"primaryKey;unique" json:"userID"`
	UserModel          UserModel `gorm:"foreignKey:UserID" json:"-"`
	OpenCommentMessage bool      `json:"openCommentMessage"` // 是否开启评论消息
	OpenDiggMessage    bool      `json:"openDiggMessage"`    // 是否开启点赞消息
	OpenPrivateChat    bool      `json:"openPrivateChat"`    // 是否开启私信消息
}
