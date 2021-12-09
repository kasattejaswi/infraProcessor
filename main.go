package main

import (
	"context"
	"flag"
)

// TODO - Write unit tests for all files
// TODO - Write dockerfile for this project

func main() {
	topic := flag.String("topic", "CPU_Topic", "Topic from which data will be read")
	brokerAddress := flag.String("brokerAddress", "localhost:9092", "Broker address in form of host:port")
	groupId := flag.String("groupId", "infraProcessorConsumer", "Consumer group ID")
	flag.Parse()
	ctx := context.Background()
	Consume(ctx, *topic, *brokerAddress, *groupId)
}
