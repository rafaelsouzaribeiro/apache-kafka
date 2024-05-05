package main

import (
	"fmt"
	"sync"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/apache-kafka/pkg/producer"
	"github.com/rafaelsouzaribeiro/apache-kafka/pkg/utils"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		Producer()
		wg.Done()
	}()

	wg.Wait()
}

func Producer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	message := utils.Message{
		Value: "Testar",
		Topic: "contact-adm-insert",
		Headers: []utils.Header{
			{
				Key:   "your-header-key1",
				Value: "your-header-value1",
			},
			{
				Key:   "your-header-key2",
				Value: "your-header-value2",
			},
		},
	}

	b, err := producer.Encode(message)

	if err != nil {
		panic(err)
	}

	produc := producer.NewProducer([]string{"springboot:9092"}, "contact-adm-insert",
		sarama.ByteEncoder(b), config, func(messages string) {
			fmt.Println("Message failure:", messages)
		})
	prod, err := produc.GetProducer()

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := (*prod).Close(); err != nil {
			panic(err)
		}
	}()

	produc.SendMessage(prod)

}
