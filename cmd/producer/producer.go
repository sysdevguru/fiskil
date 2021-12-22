package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sysdevguru/fiskil/config"
	"github.com/sysdevguru/fiskil/models"
)

type Severity int

const (
	Debug Severity = iota
	Info
	Warn
	Error
	Fatal
)

func (s Severity) String() string {
	return []string{"debug", "info", "warn", "error", "fatal"}[s]
}

func generateLog(i int) []byte {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	log := models.Log{
		ServiceName: fmt.Sprintf("service_%d", i),
		Payload:     "log content",
		Severity:    Severity(r.Intn(100) % 5).String(),
	}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(log)

	return reqBodyBytes.Bytes()
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.Kafka.URL, cfg.Kafka.Topic, cfg.Kafka.Partition)
	if err != nil {
		log.Fatalln("failed to dial leader:", err)
	}

	logsCount := 5
	iteration := 2000

	conn.SetWriteDeadline(time.Now().Add(60 * time.Second))
	for s := 0; s < iteration; s++ {
		for i := 0; i < logsCount; i++ {
			logData := generateLog(i)
			_, err = conn.WriteMessages(
				kafka.Message{Value: logData},
			)

			if err != nil {
				log.Fatalln("failed to write messages:", err)
			}
		}
	}

	log.Printf("wrote %d logs...", logsCount*iteration)

	if err := conn.Close(); err != nil {
		log.Fatalln("failed to close writer:", err)
	}

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
}
