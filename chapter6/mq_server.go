package chapter6

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
)

type MqServer struct {
	conn *amqp.Connection
	Ch   *amqp.Channel
}

var once sync.Once
var mqs *MqServer

func NewMqServer() *MqServer {
	url := "amqp://root:123123@127.0.0.1:5672/chang"
	once.Do(func() {
		conn, err := amqp.Dial(url)
		if err != nil {
			panic(err)
		}

		ch, err := conn.Channel()
		if err != nil {
			panic(err)
		}

		ch.Qos(1, 0, false)

		mqs = &MqServer{
			conn: conn,
			Ch:   ch,
		}
	})

	return mqs
}

// 声明队列
func (mq *MqServer) DeclareQueue(queueName string) error {
	_, err := mq.Ch.QueueDeclare(queueName, false, false, false, false, nil)
	return err
}

// 声明交换机
func (mq *MqServer) DeclareExchange(exchangeName string) error {
	err := mq.Ch.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	return err
}

// 绑定队列
func (mq *MqServer) BindQueue(queueName, exchangeName, routeKey string) error {
	err := mq.Ch.QueueBind(queueName, routeKey, exchangeName, false, nil)
	return err
}

// 发送消息
func (mqs *MqServer) SendMessage(exchangeName, routeKey, message string) error {
	return mqs.Ch.Publish(exchangeName, routeKey, false, false, amqp.Publishing{
		Body: []byte(message),
	})
}

// 发送带有效期的消息
func (mq *MqServer) SendMsgExp(exchange, route string, msg []byte, expire int) error {
	return mq.Ch.Publish(exchange, route, false, false, amqp.Publishing{
		Body:       msg,
		Expiration: fmt.Sprintf("%d", expire*1000),
	})
}

// 接收消息
func (mq *MqServer) ReceiveMsg(queue string, handler func([]byte) error) {

	msgs, err := mq.Ch.Consume(queue, "", false, false, false, false, nil)
	if err != nil {
		log.Println("Failed to receive message err:", err.Error())
		return
	}

	for msg := range msgs {
		err = handler(msg.Body)
		if err != nil {
			log.Println("Failed to handle message err:", err.Error())
		}
		msg.Ack(false)
	}
}

// 多个消费者接收消息
func (mq *MqServer) MulReceiveMsg(queue string, handler func([]byte) error, quantity int) {

	for i := 0; i < quantity; i++ {
		go mq.ReceiveMsg(queue, handler)
	}

	quit := make(chan struct{})
	<-quit
}

// 声明延迟队列
func (mq *MqServer) DeclareDelayQueue(exchangeName, queueName, toQueue string) error {
	_, err := mq.Ch.QueueDeclare(queueName, false, false, false, false, amqp.Table{
		"x-dead-letter-exchange":    exchangeName,
		"x-dead-letter-routing-key": toQueue,
	})
	if err != nil {
		return err
	}
	err = mq.Ch.QueueBind(queueName, queueName, exchangeName, false, nil)
	return err
}
