package producer

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type Producer interface {
	
}


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
	err := p.writer.Close()
	if err != nil {
		log.Fatal("failed to close client:", err)
	}
}

func NewProducer(broker string, topic string) *ProducerClient {
	w := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:   topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &ProducerClient{
		writer: w,
	}
	
}