package topic

import (
	"log"
	"strings"

	"github.com/Shopify/sarama"

	"topic/pkg/kafkautil"
)

type Inputs struct {
	Urls         string
	Name        string
	Replication int16
	Partitions  int32
}

func Create(i *Inputs) {
	c := sarama.NewConfig()
	c.Version = kafkautil.PromptVersion()

	ca, err := sarama.NewClusterAdmin(strings.Split(i.Urls, ","), c)
	if err != nil {
		log.Println(err)
		return
	}

	d := sarama.TopicDetail{NumPartitions: i.Partitions, ReplicationFactor: i.Replication}
	err = ca.CreateTopic(i.Name, &d, false)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Kafka topic created successfully!")
}
