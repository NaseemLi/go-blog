package chatmsg

type TextMsg struct {
	Content string `json:"content"` // 消息内容
}

type ImageMsg struct {
	Href string `json:"href"` // 图片地址
	Src  string `json:"src"`
}

type MarkDownMsg struct {
	Content string `json:"content"` // 消息内容
}

type ChatMsg struct {
	TextMsg     *TextMsg     `json:"textMsg,omitempty"`
	ImageMsg    *ImageMsg    `json:"imageMsg,omitempty"`
	MarkDownMsg *MarkDownMsg `json:"markDownMsg,omitempty"`
}
