package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/skhanal5/kafka-tryout/internal/config"
	"github.com/skhanal5/kafka-tryout/internal/consumer"
)

func main() {
	cfg := config.NewConfig()
	c := consumer.NewConsumer(cfg.KafkaBroker, cfg.KafkaTopic)

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

    // Start message production loop
    go c.ReadMessage(context.Background())

    // Wait for shutdown signal
    <-stop
    log.Println("shutting down consumer client...")
    c.Shutdown()
}