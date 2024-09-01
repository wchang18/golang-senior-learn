package chapter6

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"sync"
	"testing"
)

func ProductSync() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.ClientID = "client01"
	config.Net.MaxOpenRequests = 10

	broker := []string{"127.0.0.1:9092", "127.0.0.1:9101"}

	producter, err := sarama.NewSyncProducer(broker, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := producter.Close(); err != nil {
			panic(err)
		}
	}()
	for i := 0; i < 10; i++ {
		part, offset, err := producter.SendMessage(&sarama.ProducerMessage{
			Topic: "test02",
			Value: sarama.StringEncoder(fmt.Sprintf("hello-world-%d", i)),
		})
		fmt.Println(i, part, offset, err)
	}

}

func TestProductSync(t *testing.T) {
	ProductSync()
}

func ProductAsync() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.ClientID = "client01"
	config.Net.MaxOpenRequests = 10

	broker := []string{"127.0.0.1:9092", "127.0.0.1:9101"}

	producter, err := sarama.NewAsyncProducer(broker, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := producter.Close(); err != nil {
			panic(err)
		}
	}()

	go func() {
		for {
			select {
			case msg := <-producter.Successes():
				fmt.Println("success", msg.Key, msg.Value, msg.Offset)
			case err = <-producter.Errors():
				fmt.Println("err", err)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		producter.Input() <- &sarama.ProducerMessage{
			Topic: "test01",
			Value: sarama.StringEncoder(fmt.Sprintf("hello-b-%d", i)),
		}
	}

	over := make(chan bool)
	<-over
}

func Consumer() {
	broker := []string{"127.0.0.1:9092", "127.0.0.1:9101"}
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(broker, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	topic := "test02"
	partList, err := consumer.Partitions(topic)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	for _, part := range partList {
		wg.Add(1)
		pc, err := consumer.ConsumePartition(topic, part, 0)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Message topic:%q partition:%d offset:%d, value:%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
			}
		}(pc)
	}
	wg.Wait()
}

func TestConsumer(t *testing.T) {
	Consumer()
}

func TestProductAsync(t *testing.T) {
	ProductAsync()
}

type ConsumerGroupHandler struct{}

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d, value:%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func ConsumerGroup() {
	broker := []string{"127.0.0.1:9092", "127.0.0.1:9101"}
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Return.Errors = true

	groupName := "my-group-2"
	group, err := sarama.NewConsumerGroup(broker, groupName, config)
	if err != nil {
		panic(err)
	}
	defer group.Close()

	go func() {
		for err := range group.Errors() {
			fmt.Println(err)
		}
	}()
	ctx := context.Background()
	topic1 := "test01"
	for {
		err := group.Consume(ctx, []string{topic1}, &ConsumerGroupHandler{})
		if err != nil {
			panic(err)
		}
	}
}

func TestConsumerGroup(t *testing.T) {
	ConsumerGroup()
}
