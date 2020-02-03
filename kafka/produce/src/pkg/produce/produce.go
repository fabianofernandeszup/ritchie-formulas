package produce

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Shopify/sarama"

	"produce/pkg/kafkautil"
)

type Inputs struct {
	Urls  string
	Topic string
}

func Run(i Inputs) {
	c := sarama.NewConfig()
	c.Version = kafkautil.PromptVersion()
	c.Producer.Return.Successes = true
	c.Producer.RequiredAcks = sarama.WaitForAll
	c.Producer.Retry.Max = 5

	p, err := sarama.NewSyncProducer(strings.Split(i.Urls, ","), c)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := p.Close(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		msg, _ := r.ReadString('\n')

		publish(i.Topic, msg, p)
	}

}

func publish(topic, msg string, producer sarama.SyncProducer) {
	m := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}

	_, _, err := producer.SendMessage(m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
