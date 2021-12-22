package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sysdevguru/fiskil/cmd/worker/command"
	"github.com/sysdevguru/fiskil/config"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("could not load configuration %s\n", err.Error())
	}
	command.Start(ctx, cfg)

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
}
