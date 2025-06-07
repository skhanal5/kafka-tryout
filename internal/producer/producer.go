package producer

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type ProducerClient struct {
	writer *kafka.Writer
}

func (p *ProducerClient) WriteMessage(ctx context.Context, message string) {
	msg := kafka.Message{
		Value: []byte(message),
	}
	err := p.writer.WriteMessages(ctx, msg)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

}

func (p *ProducerClient) Shutdown() {
	if err := p.writer.Close(); err != nil {
		log.Println("failed to close producer:", err)
	}
}

func NewProducer(broker string, topic string) *ProducerClient {
	w := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &ProducerClient{
		writer: w,
	}

}
