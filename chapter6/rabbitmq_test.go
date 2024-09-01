package chapter6

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"testing"
	"time"
)

func TestConn(t *testing.T) {
	mq := NewMqServer()
	fmt.Println(mq)
}

func TestDeclareQueue(t *testing.T) {
	mq := NewMqServer()
	//声明交换机
	err := mq.Ch.ExchangeDeclare("default_exchange", "direct", true, false, false, false, nil)
	if err != nil {
		t.Fatal(err)
	}

	//声明队列
	_, err = mq.Ch.QueueDeclare("test_queue_1", true, false, false, false, nil)
	if err != nil {
		t.Fatal(err)
	}

	//绑定队列到交换机
	err = mq.Ch.QueueBind("test_queue_1", "test_queue_1", "default_exchange", false, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPublish(t *testing.T) {
	mq := NewMqServer()
	err := mq.Ch.Publish("default_exchange", "test_queue_1", false, false, amqp.Publishing{
		Body: []byte("hello-001"),
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestPublishMany(t *testing.T) {
	mq := NewMqServer()
	for i := 1; i <= 100; i++ {
		message := fmt.Sprintf("hello-%d", i)
		err := mq.Ch.Publish("default_exchange", "test_queue_1", false, false, amqp.Publishing{
			Body: []byte(message),
		})
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestConsume(t *testing.T) {
	mq := NewMqServer()
	msgs, err := mq.Ch.Consume("test_queue_1", "", false, false, false, false, nil)
	if err != nil {
		t.Fatal(err)
	}

	for msg := range msgs {
		t.Log(string(msg.Body))
		msg.Ack(false)
	}
}

func TestConsumerMany(t *testing.T) {
	mq := NewMqServer()
	for i := 0; i < 3; i++ {
		go func(n int) {
			consumer := fmt.Sprintf("consumer-id:%d", n)
			fmt.Println(consumer)
			msgs, err := mq.Ch.Consume("test_queue_1", consumer, false, false, false, false, nil)
			if err != nil {
				t.Fatal(err)
			}

			for msg := range msgs {
				t.Log("consumer", n, string(msg.Body))
				time.Sleep(time.Second)
				msg.Ack(false)
			}
		}(i)
	}
	quit := make(chan struct{})
	<-quit
}

func TestDeclare(t *testing.T) {
	mq := NewMqServer()
	for i := 1; i <= 3; i++ {
		name := fmt.Sprintf("order_status_%d", i)
		_, err := mq.Ch.QueueDeclare(name, true, false, false, false, nil)
		if err != nil {
			t.Fatal(err)
		}

		mq.Ch.QueueBind(name, "order_status", "default_exchange", false, nil)
	}
}

func TestPublish2(t *testing.T) {
	mq := NewMqServer()
	err := mq.Ch.Publish("default_exchange", "order_status", false, false, amqp.Publishing{
		Body: []byte("hello-001"),
	})
	if err != nil {
		t.Fatal(err)
	}
}
