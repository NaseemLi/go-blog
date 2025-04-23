package chatmsgtypeenum

type MsgType int8

const (
	TextMsgType MsgType = iota + 1
	ImageMsgType
	MarkDownMsgType
)
