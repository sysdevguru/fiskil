package policy

import (
	"fmt"

	gocraft "github.com/gocraft/work"
	"github.com/sysdevguru/fiskil/config"
)

var jobOptions = map[string]gocraft.JobOptions{
	config.TaskParser:  ExponentialRetryPolicy,
	config.TaskUpdater: ExponentialRetryPolicy,
	config.TaskFailure: ExponentialRetryPolicy,
}

func MustGetJobOption(name string) gocraft.JobOptions {
	val, ok := jobOptions[name]
	if !ok {
		panic(fmt.Sprintf("JobOption %s does not exist", name))
	}

	return val
}
