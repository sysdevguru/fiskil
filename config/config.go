package config

import (
	"github.com/kelseyhightower/envconfig"
)

const (
	envPrefix = "settings"
)

const (
	TaskParser  = "task.parser"
	TaskUpdater = "task.updater"
	TaskFailure = "task.failure"
)

const (
	LoggerPool  string = "logger_pool"
	FailurePool string = "failure_pool"
)

// App contains configuration for app.
type App struct {
	BatchRecordCount int    `envconfig:"APP_BATCH_RECORD_COUNT" default:"5000"`
	FlushInterval    string `envconfig:"APP_FLUSH_INTERVAL" required:"true"`
}

// Database contains configuration for database.
type Database struct {
	URL                string `envconfig:"DATABASE_URL" required:"true"`
	LogLevel           string `envconfig:"DATABASE_LOG_LEVEL" default:"warn"`
	MaxOpenConnections int    `envconfig:"DATABASE_MAX_OPEN_CONNECTIONS" default:"10"`
}

// Kafka contains configuration for Kafka.
type Kafka struct {
	URL       string `envconfig:"KAFKA_URL" default:"kafka:9092"`
	Topic     string `envconfig:"KAFKA_TOPIC" default:"fiskil"`
	Partition int    `envconfig:"KAFKA_PARTITION" default:"0"`
	ReadMin   int    `envconfig:"KAFKA_READ_MIN" default:"100000"`
	ReadMax   int    `envconfig:"KAFKA_READ_MAX" default:"10000000"`
}

// Task contains configurations for Gocraft workers.
type Task struct {
	NumWorkers uint   `envconfig:"TASK_NUM_WORKERS" default:"1"`
	RedisURL   string `envconfig:"TASK_REDIS_URL" default:""`
	MaxActive  int    `envconfig:"TASK_MAX_ACTIVE" default:"1"`
	MaxIdle    int    `envconfig:"TASK_MAX_IDLE" default:"1"`
}

// Config is the global config struct.
type Config struct {
	App      App
	Database Database
	Kafka    Kafka
	Task     Task
}

// Load configuration from environment.
func Load() (Config, error) {
	config := Config{}

	if err := envconfig.Process(envPrefix, &config); err != nil {
		return config, err
	}

	return config, nil
}
