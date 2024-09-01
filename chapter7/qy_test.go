package chapter7

import (
	"fmt"
	"testing"
)

func TestSendTxt(t *testing.T) {
	qy := NewQyWechatMessage("bf0c6df7-396c-4a3a-9f84-f0c633d3f9e6")
	msg := "env:pre\n"
	msg += "用户数：100\n"
	msg += "订单数：1000"
	err := qy.SendTextMessage(msg)
	t.Log(err)
}

func TestSendMd(t *testing.T) {
	qy := NewQyWechatMessage("bf0c6df7-396c-4a3a-9f84-f0c633d3f9e6")
	msg := fmt.Sprintf("## env: %s\n", "pre")
	msg += fmt.Sprintf("新增订单：<font color=\"info\">%d</font>\n", 200)
	msg += fmt.Sprintf("流失用户数：<font color=\"warning\">%d</font>", 30)
	err := qy.SendMarkdownMessage(msg)
	t.Log(err)
}
