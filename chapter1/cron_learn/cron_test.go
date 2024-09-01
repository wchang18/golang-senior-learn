package cron_learn

import (
	"github.com/robfig/cron/v3"
	"log"
	"testing"
)

func TestCronM(t *testing.T) {
	c := cron.New()
	c.AddFunc("* * * * *", func() {
		log.Println("hello")
	})

	c.Run()
}

func TestCronS(t *testing.T) {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/5 * * * * *", func() {
		log.Println("hello")
	})

	c.Run()
}

func TestCron2(t *testing.T) {
	c := cron.New()
	// 每5秒执行一次
	c.AddFunc("@every 5s", func() {
		log.Println("hello1")
	})
	// 每2分钟执行一次
	c.AddFunc("@every 2m", func() {
		log.Println("hello2")
	})
	// 每小时执行一次
	c.AddFunc("@hourly", func() {
		log.Println("hello3")
	})

	c.Run()
}
