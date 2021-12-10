package main

// Import of packages
import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

// Consume function which consumes the data from kafka queue
func Consume(ctx context.Context, topic string, brokerAddress string, groupId string, outFilePath string, timeout int) {

	// Assigning kafka values to ReaderConfig type and passing to NewReader function
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: groupId,
	})

	// Creating a channel for quitting listening of consumer after certain time
	quit := make(chan bool)

	// Go routine for listenening to queue
	go func() {
		fmt.Println("Listener started on", brokerAddress)
		for {
			select {
			// In case channel receives data, stop the execution of go routine
			case <-quit:
				fmt.Println("Stopping listener")
				return
			default:
				// Passing context to ReadMessage function for async operation and getting messages and errors
				msg, err := r.ReadMessage(ctx)

				// Terminating in case any errors are encountered
				if err != nil {
					panic(err)
				}
				if msg.Value != nil {
					// Sending the data to the file for writing
					WriteToFile(outFilePath, string(msg.Value))
				}
			}
		}
	}()
	// Run the listener forever if timeout is less than 0
	if timeout < 0 {
		for {
			time.Sleep(time.Duration(time.Second * 100))
		}
	} else {
		time.Sleep(time.Second * time.Duration(timeout))
		quit <- true
	}

}
