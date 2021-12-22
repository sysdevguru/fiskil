package task

import (
	gocraft "github.com/gocraft/work"
	"github.com/sysdevguru/fiskil/config"
	"github.com/sysdevguru/fiskil/task/handler"
	"github.com/sysdevguru/fiskil/task/policy"
)

func (s *Service) createLoggerPool(cfg config.Config, deps *handler.ServiceDependencies) error {
	pool := gocraft.NewWorkerPool(
		handler.Context{},
		cfg.Task.NumWorkers,
		config.LoggerPool,
		s.redis,
	)

	middlewares, err := handler.EventMiddlewares(cfg, deps)
	if err != nil {
		return err
	}

	for _, mdw := range middlewares {
		pool.Middleware(mdw)
	}

	pool.PeriodicallyEnqueue(cfg.App.FlushInterval, config.TaskParser)

	pool.JobWithOptions(
		config.TaskParser,
		policy.MustGetJobOption(config.TaskParser),
		(*handler.Context).ProcessNewLogParseTask,
	)

	pool.JobWithOptions(
		config.TaskUpdater,
		policy.MustGetJobOption(config.TaskUpdater),
		(*handler.Context).ProcessNewLogUpdateTask,
	)

	s.pools = append(s.pools, pool)

	return nil
}

func (s *Service) failurePool(cfg config.Config, deps *handler.ServiceDependencies) error {
	pool := gocraft.NewWorkerPool(
		handler.Context{},
		cfg.Task.NumWorkers,
		config.FailurePool,
		s.redis,
	)

	middlewares, err := handler.EventMiddlewares(cfg, deps)
	if err != nil {
		return err
	}

	for _, mdw := range middlewares {
		pool.Middleware(mdw)
	}

	pool.JobWithOptions(
		config.TaskFailure,
		policy.MustGetJobOption(config.TaskFailure),
		(*handler.Context).ProcessNewFailureTask,
	)

	s.pools = append(s.pools, pool)

	return nil
}
