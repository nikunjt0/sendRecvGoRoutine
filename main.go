package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatal("Failed to create producer: ", err)
	}
	defer producer.Close()

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatal("Failed to create consumer: ", err)
	}
	defer consumer.Close()

	go gSend(producer)
	go gRecv(consumer)

	// Wait indefinitely to keep the goroutines running
	select {}
}

func gSend(producer sarama.SyncProducer) {
	for {
		data := struct {
			Datetime string `json:"datetime"`
			Data     int    `json:"data"`
		}{
			Datetime: time.Now().UTC().Format(time.RFC3339),
			Data:     rand.Intn(100),
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println("Failed to marshal JSON data: ", err)
			continue
		}

		message := &sarama.ProducerMessage{
			Topic: "test_topic",
			Value: sarama.StringEncoder(jsonData),
		}

		_, _, err = producer.SendMessage(message)
		if err != nil {
			log.Println("Failed to send message: ", err)
		}

		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
}

func gRecv(consumer sarama.Consumer) {
    partitionConsumer, err := consumer.ConsumePartition("test_topic", 0, sarama.OffsetOldest)
    if err != nil {
        log.Fatal("Failed to start consumer: ", err)
    }
    
    done := make(chan struct{})
    defer close(done)
    
    go func() {
        for {
            select {
            case msg := <-partitionConsumer.Messages():
                var data struct {
                    Datetime string `json:"datetime"`
                    Data     int    `json:"data"`
                }
                err := json.Unmarshal(msg.Value, &data)
                if err != nil {
                    log.Println("Failed to unmarshal JSON data: ", err)
                    continue
                }
                
                fmt.Printf("\nReceived message:\nDatetime: %s\nData: %d\n\n", data.Datetime, data.Data)
            
            case err := <-partitionConsumer.Errors():
                log.Println("Consumer error: ", err.Err)
            
            case <-done:
                return
            }
        }
    }()
    
    // Wait indefinitely to keep the goroutine running
    select {}
}

