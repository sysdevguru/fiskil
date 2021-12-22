package policy

import (
	gocraft "github.com/gocraft/work"
)

// ExponentialRetryPolicy tries 8 times by exponentially increasing backoff
// It's be 1s, 4s, 9s, 16s, 25, 36, 49, 64s ~> 3:24
// After 8 times, it will drop the payload.
var ExponentialRetryPolicy = gocraft.JobOptions{
	Backoff: func(job *gocraft.Job) int64 {
		return job.Fails * job.Fails
	},
	MaxConcurrency: 1,
	MaxFails:       1,
	Priority:       10000,
	SkipDead:       true,
}
