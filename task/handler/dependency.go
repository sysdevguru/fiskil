package handler

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"github.com/sysdevguru/fiskil/config"
	"github.com/sysdevguru/fiskil/dependency"
	"github.com/sysdevguru/fiskil/task/enqueuer"
	"github.com/sysdevguru/fiskil/usecase/dcl"
	"gorm.io/gorm"
)

type ServiceDependencies struct {
	enqueuer enqueuer.Enqueuer
	logger   *dcl.UseCase
}

func (s *ServiceDependencies) StartScheduler() error {
	return s.enqueuer.AddNewLogParseTask()
}

func NewServiceDependencies(cfg config.Config, db *gorm.DB, redisPool *redis.Pool) (*ServiceDependencies, error) {
	enqueuer, err := enqueuer.NewEnqueuer(cfg, redisPool)
	if err != nil {
		return nil, errors.Wrap(err, "Enqueuer failed to initialise")
	}
	logger := dcl.New(
		config.TaskParser,
		dependency.NewRepo(db),
	)

	return &ServiceDependencies{
		enqueuer,
		logger,
	}, nil
}
