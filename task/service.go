package task

import (
	"fmt"

	gocraft "github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"github.com/sysdevguru/fiskil/config"
	"github.com/sysdevguru/fiskil/dependency"
	"github.com/sysdevguru/fiskil/task/handler"
	"gorm.io/gorm"
)

// Service is a task service which can include one or many pools of workers.
type Service struct {
	pools []*gocraft.WorkerPool
	db    *gorm.DB
	deps  *handler.ServiceDependencies
	redis *redis.Pool
}

// NewService returns a newly configured Service.
func NewService(cfg config.Config) (*Service, error) {
	service := &Service{}

	if err := service.configRedis(cfg); err != nil {
		return &Service{}, err
	}

	if err := service.configDB(cfg); err != nil {
		return &Service{}, err
	}

	deps, err := handler.NewServiceDependencies(cfg, service.db, service.redis)
	if err != nil {
		return &Service{}, err
	}

	service.deps = deps

	if err := service.createLoggerPool(cfg, deps); err != nil {
		return &Service{}, err
	}

	if err := service.failurePool(cfg, deps); err != nil {
		return &Service{}, err
	}

	return service, nil
}

// Start starts the workers in each pool.
func (s *Service) Start() {
	for _, p := range s.pools {
		p.Start()
	}

	err := s.deps.StartScheduler()
	if err != nil {
		fmt.Println("failed to fire first batch insert:", err)
	}
}

// Stop stops the workers in each pool.
func (s *Service) Stop() {
	for _, p := range s.pools {
		p.Stop()
	}
	dependency.CloseDatabaseConnection(s.db)
}

func (s *Service) configRedis(cfg config.Config) error {
	redisPool, err := dependency.NewRedisPool(cfg)
	if err != nil {
		return errors.Wrap(err, "Redis failed to connect")
	}

	s.redis = redisPool
	return nil
}

func (s *Service) configDB(cfg config.Config) error {
	db, err := dependency.NewGormWithPostgres(cfg)
	if err != nil {
		return errors.Wrap(err, "Gorm with Postgres failed to connect")
	}

	s.db = db
	return nil
}
