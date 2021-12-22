package handler

import (
	"context"

	"github.com/segmentio/kafka-go"
)

// Context holds all the data for each task
type Context struct {
	deps *ServiceDependencies
	ctx  context.Context
	conn *kafka.Conn
}
