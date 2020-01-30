package topic

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"

	"topic/pkg/kafkautil"
)

type Inputs struct {
	Url string
}

func List(i Inputs) {
	c := sarama.NewConfig()
	c.Version = kafkautil.PromptVersion()

	ca, err := sarama.NewClusterAdmin([]string{i.Url}, c)
	if err != nil {
		log.Println(err)
		return
	}

	tt, _ := ca.ListTopics()

	for k, _ := range tt {
		fmt.Println(k)
	}
}
