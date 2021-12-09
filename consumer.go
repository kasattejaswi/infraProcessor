package main

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func Consume(ctx context.Context, topic string, brokerAddress string, groupId string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: groupId,
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic(err)
		}
		WriteToFile("cpuData.txt", string(msg.Value))
	}
}
