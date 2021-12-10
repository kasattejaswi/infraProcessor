package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
)

// This test will only work when kafka service is running properly and a producer is writing data to it
// Time interval of producing data should be less than 10 seconds
func TestConsume(t *testing.T) {
	fmt.Println("Running consumer test. It will take almost 30 seconds to run")
	ctx := context.Background()
	topic := "CPU_Topic"
	groupId := "infraProcessorConsumer"
	outFile := "outfile13289uwjje89wdsjner"
	brokerAddress := "localhost:9092"
	timeout := 30
	Consume(ctx, topic, brokerAddress, groupId, outFile, timeout)

	// Check if file is created after consumer ran for 30 seconds
	if _, err := os.Stat(outFile); errors.Is(err, os.ErrNotExist) {
		t.Error("File is not created after consuming kafka queue")
	}
	// Deleting temp file after testing is completed
	e := os.Remove(outFile)
	if e != nil {
		panic(e)
	}
}
