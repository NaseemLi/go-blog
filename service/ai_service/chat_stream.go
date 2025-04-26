package aiservice

import (
	"bufio"
	_ "embed"
	"encoding/json"

	"github.com/sirupsen/logrus"
)

//go:embed chat_stream.prompt
var chatStreamPrompt string

// 用于流式响应的结构体
type StreamChoice struct {
	Index        int          `json:"index"`
	Delta        DeltaMessage `json:"delta"`
	FinishReason string       `json:"finish_reason"`
}

type DeltaMessage struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

// 流式响应的完整结构
type ChatStreamResponse struct {
	ID                string         `json:"id"`
	Object            string         `json:"object"`
	Created           int64          `json:"created"`
	Model             string         `json:"model"`
	Choices           []StreamChoice `json:"choices"`
	SystemFingerprint *string        `json:"system_fingerprint,omitempty"`
}

func ChatStream(content string, params string) (msgChan chan string, err error) {
	msgChan = make(chan string)
	r := Request{
		Model: "gpt-3.5-turbo",
		Messages: []Messages{
			{
				Role:    "system",
				Content: chatStreamPrompt + params,
			},
			{
				Role:    "user",
				Content: content,
			},
		},
		Stream: true,
	}
	res, err := BaseRequest(r)
	if err != nil {
		return nil, err // 这里一定要返回错误，不然外层判断是空 chan 就会卡死
	}

	scanner := bufio.NewScanner(res.Body)
	scanner.Split(bufio.ScanLines)

	go func() {
		defer res.Body.Close()
		defer close(msgChan)

		for scanner.Scan() {
			text := scanner.Text()
			if text == "" {
				continue
			}
			if len(text) < 6 || text[:6] != "data: " {
				continue
			}
			data := text[6:]
			if data == "[DONE]" {
				return
			}
			var item ChatStreamResponse
			err := json.Unmarshal([]byte(data), &item)
			if err != nil {
				logrus.Errorf("解析失败 %v %s", err, data)
				continue
			}
			if len(item.Choices) > 0 {
				msgChan <- item.Choices[0].Delta.Content
			}
		}
	}()

	return msgChan, nil
}
