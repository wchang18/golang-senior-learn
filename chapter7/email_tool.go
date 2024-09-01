package chapter7

import "gopkg.in/gomail.v2"

type EmailTool struct {
	Host   string
	Port   int
	User   string
	Pass   string
	Dialer *gomail.Dialer
}

func NewEmailTool(host string, port int, user string, pass string) *EmailTool {
	return &EmailTool{
		Host:   host,
		Port:   port,
		User:   user,
		Pass:   pass,
		Dialer: gomail.NewDialer(host, port, user, pass),
	}
}

func (e *EmailTool) Send(to []string, subject string, body string, files ...string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.User)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	for _, file := range files {
		m.Attach(file)
	}
	return e.Dialer.DialAndSend(m)
}
