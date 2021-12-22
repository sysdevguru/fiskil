package handler

import (
	"context"
	"time"

	gocraft "github.com/gocraft/work"
	"github.com/segmentio/kafka-go"
	"github.com/sysdevguru/fiskil/config"
)

type MiddlewareFn func(*Context, *gocraft.Job, gocraft.NextMiddlewareFunc) error

func EventMiddlewares(cfg config.Config, deps *ServiceDependencies) ([]MiddlewareFn, error) {
	return []MiddlewareFn{
		configJobDependencies(deps, cfg),
	}, nil
}

func configJobDependencies(deps *ServiceDependencies, cfg config.Config) func(*Context, *gocraft.Job, gocraft.NextMiddlewareFunc) error {
	return func(c *Context, job *gocraft.Job, next gocraft.NextMiddlewareFunc) error {
		c.ctx = context.Background()
		c.deps = deps
		c.conn = configKafkaConn(cfg)
		return next()
	}
}

func configKafkaConn(cfg config.Config) *kafka.Conn {
	conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.Kafka.URL, cfg.Kafka.Topic, cfg.Kafka.Partition)
	if err != nil {
		return nil
	}
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	return conn
}
