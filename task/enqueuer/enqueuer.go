package enqueuer

import (
	"github.com/gocraft/work"
	gocraft "github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"github.com/sysdevguru/fiskil/config"
	"github.com/sysdevguru/fiskil/models"
)

type Enqueuer interface {
	AddNewLogParseTask() error
	AddNewLogUpdateTask(logs *[]models.Log, severities *[]models.Severity) error
}

type GocraftEnqueuer struct {
	loggingEnqueuer *work.Enqueuer
	failureEnqueuer *work.Enqueuer
}

func (enq *GocraftEnqueuer) AddNewLogParseTask() error {
	if _, err := enq.loggingEnqueuer.Enqueue(
		config.TaskParser,
		gocraft.Q{},
	); err != nil {
		return errors.Wrap(err, "unable to enqueue AddNewLogParseTask")
	}

	return nil
}

func (enq *GocraftEnqueuer) AddNewLogUpdateTask(logs *[]models.Log, severities *[]models.Severity) error {
	if _, err := enq.loggingEnqueuer.Enqueue(
		config.TaskUpdater,
		gocraft.Q{
			"logs":       logs,
			"severities": severities,
		},
	); err != nil {
		return errors.Wrap(err, "unable to enqueue AddNewLogInsertTask")
	}

	return nil
}

func (enq *GocraftEnqueuer) AddFailedJob(job *gocraft.Job) error {
	if _, err := enq.failureEnqueuer.Enqueue(
		config.TaskFailure,
		gocraft.Q{
			"originalJobName": job.Name,
		},
	); err != nil {
		return errors.Wrap(err, "unable to enqueue AddFailureTask")
	}

	return nil
}

func NewEnqueuer(cfg config.Config, redisPool *redis.Pool) (*GocraftEnqueuer, error) {
	return &GocraftEnqueuer{
		loggingEnqueuer: gocraft.NewEnqueuer(config.LoggerPool, redisPool),
		failureEnqueuer: gocraft.NewEnqueuer(config.FailurePool, redisPool),
	}, nil
}
