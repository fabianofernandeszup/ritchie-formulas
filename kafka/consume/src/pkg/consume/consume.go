package consume

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
	"github.com/hashicorp/go-uuid"

	"consume/pkg/kafkautil"
)

const ritchieGroupFormat = "ritchie_consumer_group_%s"

type Inputs struct {
	Urls          string
	Topic         string
	FromBeginning bool
}

func Run(i Inputs) {
	c := sarama.NewConfig()
	c.Version = kafkautil.PromptVersion()
	c.Consumer.Return.Errors = true

	if i.FromBeginning {
		c.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	consumer := Consumer{
		ready: make(chan bool),
	}

	ctx, cancel := context.WithCancel(context.Background())
	uuid, _ := uuid.GenerateUUID()
	client, err := sarama.NewConsumerGroup(strings.Split(i.Urls, ","), fmt.Sprintf(ritchieGroupFormat, uuid), c)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			if err := client.Consume(ctx, []string{i.Topic}, &consumer); err != nil {
				fmt.Println(fmt.Sprintf("Error from consumer: %v", err))
				os.Exit(1)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // Await till the consumer has been set up
	fmt.Println("Ritchie consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		fmt.Println("Terminating: context cancelled")
	case <-sigterm:
		fmt.Println("Terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

// Consumer represents a Ritchie consumer group consumer
type Consumer struct {
	ready chan bool
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Println(fmt.Sprintf("Consumed message from topic (%s) at %s: \n %s \n", message.Topic, message.Timestamp.Format("2006-01-02T15:04:05.0000"), string(message.Value)))
		session.MarkMessage(message, "")
	}

	return nil
}
