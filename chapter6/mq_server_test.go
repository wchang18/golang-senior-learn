package chapter6

import (
	"fmt"
	"testing"
	"time"
)

func TestDeclareDelayQueue(t *testing.T) {
	mq := NewMqServer()
	//声明数据处理队列
	mq.DeclareExchange("delay_exchange")
	mq.DeclareQueue("report_data")
	mq.BindQueue("report_data", "delay_exchange", "report_data")

	mq.DeclareDelayQueue("delay_exchange", "wait_time_20s", "report_data")
}

func TestSend(t *testing.T) {
	mq := NewMqServer()
	for i := 0; i < 100; i++ {
		mq.SendMsgExp("delay_exchange", "wait_time_20s", []byte("hello-"+fmt.Sprint(i)), 20)
	}
}

func TestMulReceiveMsg(t *testing.T) {
	mq := NewMqServer()
	mq.MulReceiveMsg("report_data", func(msg []byte) error {
		fmt.Println(string(msg))
		time.Sleep(time.Second)
		return nil
	}, 3)
}
