package aiservice

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"goblog/global" // 替代某些 io/ioutil 函数
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model    string     `json:"model"`
	Messages []Messages `json:"messages"`
	Stream   bool       `json:"stream"`
}

type ChatResponse struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
	SystemFingerprint *string  `json:"system_fingerprint"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens            int          `json:"prompt_tokens"`
	CompletionTokens        int          `json:"completion_tokens"`
	TotalTokens             int          `json:"total_tokens"`
	CompletionTokensDetails TokenDetails `json:"completion_tokens_details"`
	PromptTokensDetails     TokenDetails `json:"prompt_tokens_details"`
}

type TokenDetails struct {
	AudioTokens     int `json:"audio_tokens"`
	ReasoningTokens int `json:"reasoning_tokens,omitempty"` // 某些字段可能缺省
	CachedTokens    int `json:"cached_tokens,omitempty"`    // 某些字段可能缺省
}

const baseUrl = "https://api.chatanywhere.tech/v1/chat/completions"

func BaseRequest(r Request) (res *http.Response, err error) {
	method := "POST"

	byteData, _ := json.Marshal(r)
	req, err := http.NewRequest(method, baseUrl, bytes.NewBuffer(byteData))
	if err != nil {
		logrus.Errorf("请求参数失败%v", err)
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", global.Config.Ai.SecretKey))
	req.Header.Add("Content-Type", "application/json")

	res, err = http.DefaultClient.Do(req)
	return

}

//go:embed chat.prompt
var chatPrompt string

func Chat(content string) (msg string, err error) {

	r := Request{
		Model: "gpt-3.5-turbo",
		Messages: []Messages{
			{
				Role:    "system",
				Content: chatPrompt,
			},
			{
				Role:    "user",
				Content: content,
			},
		},
		Stream: false,
	}
	res, err := BaseRequest(r)
	if err != nil {
		return
	}
	body, _ := ioutil.ReadAll(res.Body)

	var response ChatResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		logrus.Errorf("解析失败%v %s", err, string(body))
		return
	}

	if len(response.Choices) == 0 {
		logrus.Warnf("AI 返回为空,content=%s,响应内容=%s", content, string(body))
		err = fmt.Errorf("AI 返回内容为空")
		return
	}

	msg = response.Choices[0].Message.Content
	return msg, nil
}
