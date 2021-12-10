package main

// Package imports
import (
	"context"
	"flag"
)

// Main function
func main() {
	// Creating commandline flags for passing values from commandline
	topic := flag.String("topic", "CPU_Topic", "Topic from which data will be read")
	brokerAddress := flag.String("brokerAddress", "localhost:9092", "Broker address in form of host:port")
	groupId := flag.String("groupId", "infraProcessorConsumer", "Consumer group ID")
	outFile := flag.String("outFile", "cpuData.txt", "File path in which output will be written")
	runTimeout := flag.Int("time", -1, "Time for which it will keep listening. For infinite run, pass -1")

	// Parsing all commandline flags
	flag.Parse()

	// Creating background context for consumer function.
	ctx := context.Background()

	// Calling consume function from consumer.go
	Consume(ctx, *topic, *brokerAddress, *groupId, *outFile, *runTimeout)
}
