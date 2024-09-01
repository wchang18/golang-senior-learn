package chapter7

import (
	"bytes"
	"fmt"
	"io"
	"k8s.io/apimachinery/pkg/util/json"
	"net/http"
)

const (
	Api = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="
)

type QyWechatMessage struct {
	Key string
}

func NewQyWechatMessage(key string) *QyWechatMessage {
	return &QyWechatMessage{
		Key: key,
	}
}

func (q *QyWechatMessage) SendTextMessage(message string) error {
	textRequest := &TextRequest{
		MsgType: "text",
		Text: Content{
			Content: message,
			//MentionedList: []string{"ACCA", "@all"},
			//MentionedMobileList: []string{},
		},
	}
	url := Api + q.Key
	return PostMessage(url, textRequest)
}

func (q *QyWechatMessage) SendMarkdownMessage(message string) error {
	textRequest := &MarkdownRequest{
		MsgType: "markdown",
		Markdown: Content{
			Content: message,
		},
	}
	url := Api + q.Key
	return PostMessage(url, textRequest)
}

func PostMessage(url string, data interface{}) error {
	b1, _ := json.Marshal(data)
	body := bytes.NewReader(b1)
	resp, err := http.Post(url, "content-type:application/json", body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	b2, _ := io.ReadAll(resp.Body)
	var response Response
	err = json.Unmarshal(b2, &response)
	if err != nil {
		return err
	}
	if response.Errcode != 0 {
		return fmt.Errorf("errcode: %d, errmsg: %s", response.Errcode, response.Errmsg)
	}
	fmt.Printf("%+v", response)
	return nil
}

type TextRequest struct {
	MsgType string  `json:"msgtype"`
	Text    Content `json:"text"`
}

type MarkdownRequest struct {
	MsgType  string  `json:"msgtype"`
	Markdown Content `json:"markdown"`
}

type Content struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

type Response struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
