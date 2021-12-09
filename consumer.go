package main

// Import of packages
import (
	"context"

	"github.com/segmentio/kafka-go"
)

// Consume function which consumes the data from kafka queue
func Consume(ctx context.Context, topic string, brokerAddress string, groupId string, outFilePath string) {

	// Assigning kafka values to ReaderConfig type and passing to NewReader function
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: groupId,
	})

	// Infinte loop to keep on reading messages
	for {
		// Passing context to ReadMessage function for async operation and getting messages and errors
		msg, err := r.ReadMessage(ctx)

		// Terminating in case any errors are encountered
		if err != nil {
			panic(err)
		}

		// Sending the data to the file for writing
		WriteToFile(outFilePath, string(msg.Value))
	}
}
