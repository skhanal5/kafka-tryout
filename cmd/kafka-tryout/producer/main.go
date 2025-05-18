package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/skhanal5/kafka-tryout/internal/config"
	"github.com/skhanal5/kafka-tryout/internal/producer"
)


func main() {
	cfg := config.NewConfig()
	c := producer.NewProducer(cfg.KafkaBroker, cfg.KafkaTopic)

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

    // Start message production loop
	go func() {
		for {
			c.WriteMessage(context.Background(), time.Now().String())
		}
	}()

    // Wait for shutdown signal
    <-stop
    c.Shutdown()
}
