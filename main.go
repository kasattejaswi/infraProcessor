package main

// Package imports
import (
	"context"
	"fmt"
	"os"
	"strconv"
)

// Main function
func main() {
	// Creating commandline flags for passing values from commandline
	// Uncomment below lines to enable command line execution

	// topic := flag.String("topic", "CPU_Topic", "Topic from which data will be read")
	// brokerAddress := flag.String("brokerAddress", "localhost:9092", "Broker address in form of host:port")
	// groupId := flag.String("groupId", "infraProcessorConsumer", "Consumer group ID")
	// outFile := flag.String("outFile", "cpuData.txt", "File path in which output will be written")
	// runTimeout := flag.Int("time", -1, "Time for which it will keep listening. For infinite run, pass -1")

	// // Parsing all commandline flags
	// flag.Parse()

	// Reading data from environment variables
	topic := os.Getenv("TOPIC")
	brokerAddress := os.Getenv("BROKER_ADDRESS")
	groupId := os.Getenv("GROUP_ID")
	outFile := os.Getenv("OUT_FILE_PATH")
	runTimeout, err := strconv.Atoi(os.Getenv("RUN_TIMEOUT"))

	// In case the string to int conversion gets failed, assign a default value of -1
	if err != nil {
		fmt.Println(err)
		fmt.Println("Setting default run timeout to infinite")
		runTimeout = -1
	}

	// Creating background context for consumer function.
	ctx := context.Background()

	// Calling consume function from consumer.go
	Consume(ctx, topic, brokerAddress, groupId, outFile, runTimeout)
}
