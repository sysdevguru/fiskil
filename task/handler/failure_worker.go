package handler

import (
	gocraft "github.com/gocraft/work"
	"github.com/sysdevguru/fiskil/config"
)

func (c *Context) ProcessNewFailureTask(job *gocraft.Job) error {
	switch jobName := job.ArgString("originalJobName"); jobName {
	case config.TaskParser, config.TaskUpdater:
		c.deps.logger.Fail(c.ctx, jobName)
	default:
	}

	return nil
}
