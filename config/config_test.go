package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	os.Setenv("APP_FLUSH_INTERVAL", "*/1 * * * *")
	os.Setenv("DATABASE_URL", "localhost:5432")

	cfg, err := Load()
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}

	t.Run("read batch insert value", func(t *testing.T) {
		if cfg.App.BatchRecordCount != 5000 {
			t.Errorf("got %d, want 5000", cfg.App.BatchRecordCount)
		}
	})

	t.Run("read batch flush interval", func(t *testing.T) {
		if cfg.App.FlushInterval != "*/1 * * * *" {
			t.Errorf("got %s, want */1 * * * *", cfg.App.FlushInterval)
		}
	})

	t.Run("read database url", func(t *testing.T) {
		if cfg.Database.URL != "localhost:5432" {
			t.Errorf("got %s, want localhost:5432", cfg.Database.URL)
		}
	})

	t.Run("read kafka url", func(t *testing.T) {
		if cfg.Kafka.URL != "kafka:9092" {
			t.Errorf("got %s, want kafka:9092", cfg.Kafka.URL)
		}
	})

	t.Run("read kafka topic", func(t *testing.T) {
		if cfg.Kafka.Topic != "fiskil" {
			t.Errorf("got %s, want fiskil", cfg.Kafka.Topic)
		}
	})

	t.Run("read kafka partition", func(t *testing.T) {
		if cfg.Kafka.Partition != 0 {
			t.Errorf("got %d, want 0", cfg.Kafka.Partition)
		}
	})

	t.Run("read kafka read min", func(t *testing.T) {
		if cfg.Kafka.ReadMin != 100000 {
			t.Errorf("got %d, want 100000", cfg.Kafka.ReadMin)
		}
	})

	t.Run("read kafka read max", func(t *testing.T) {
		if cfg.Kafka.ReadMax != 10000000 {
			t.Errorf("got %d, want 10000000", cfg.Kafka.ReadMax)
		}
	})
}
