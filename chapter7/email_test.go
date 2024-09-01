package chapter7

import (
	"fmt"
	"log"
	"net/smtp"
	"testing"
)

func TestEmail(t *testing.T) {
	// 邮件服务器配置
	smtpServer := "smtp.163.com:25" // 也可以使用端口465或587
	auth := smtp.PlainAuth("", "wchang18@163.com", "KHMBRPHVBUUWTRGH", "smtp.163.com")

	// 邮件内容
	from := "wchang18@163.com"
	to := []string{"1771164357@qq.com"}
	subject := "Hello Gary"
	body := "This is a test email."

	// 创建邮件消息
	message := fmt.Sprintf("To:%s\r\n"+
		"Subject: %s\r\n"+
		"\r\n "+
		"%s\r\n", to[0], subject, body)

	// 发送邮件
	err := smtp.SendMail(smtpServer, auth, from, to, []byte(message))
	if err != nil {
		log.Fatalf("Failed to send email, error: %v", err)
	}

	log.Println("Email send successfully!")
}

func TestGoEmail(t *testing.T) {
	emailTool := NewEmailTool("smtp.163.com", 25, "wchang18@163.com", "KHMBRPHVBUUWTRGH")
	emailTool.Send([]string{"1771164357@qq.com"}, "2024年8月的报表文件", "<p>请下载文件</p>", "./202408.txt")
	t.Log("Email send successfully!")
}
