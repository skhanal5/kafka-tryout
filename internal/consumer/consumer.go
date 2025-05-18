package consumer

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type ConsumerClient struct {
	reader *kafka.Reader
}

func (c *ConsumerClient) ReadMessage(ctx context.Context) {
	for {
		m, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
}

func (c *ConsumerClient) Shutdown() {
	err := c.reader.Close()
	if err != nil {
		log.Fatal("failed to close reader client:", err)
	}
}

func NewProducer(broker string, topic string) *ConsumerClient {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:    []string{broker},
		Topic:   topic,
		Partition: 0,
		MaxBytes:  10e6,
	})
	return &ConsumerClient{
		reader: r,
	}
	
}