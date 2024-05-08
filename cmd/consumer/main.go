package main

import (
	"github.com/rafaelsouzaribeiro/apache-kafka/pkg/kafka/consumer"
	"github.com/rafaelsouzaribeiro/apache-kafka/pkg/utils"
)

func main() {
	data := utils.Message{
		Topics:    []string{"contact-adm-insert"},
		Topic:     "contact-adm-insert",
		GroupID:   "contacts",
		Partition: 0,
		Offset:    -1,
	}
	go consumer.Consumer(&[]string{"springboot:9092"}, &data)
	go consumer.ListenPartition(&[]string{"springboot:9092"}, &data)
	select {}

}
