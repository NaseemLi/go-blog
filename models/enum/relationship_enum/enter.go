package relationshipenum

type Relation int8

const (
	// 陌生人，谁也没关注谁
	RelationStranger Relation = iota
	// 我关注了对方（单向关注）
	RelationFocus
	// 对方关注了我（我没回关）
	RelationFans
	// 好友，互相关注
	RelationFriend
)
