package dependency

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/sysdevguru/fiskil/config"
)

func NewRedisPool(c config.Config) (*redis.Pool, error) {
	if c.Task.RedisURL == "" {
		return nil, nil
	}

	pool := &redis.Pool{
		MaxActive: c.Task.MaxActive,
		MaxIdle:   c.Task.MaxIdle,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(c.Task.RedisURL,
				// Read timeout should be greater than ping period.
				redis.DialReadTimeout(45*time.Second),
				redis.DialWriteTimeout(10*time.Second),
			)
		},
	}

	conn := pool.Get()
	defer conn.Close()

	if err := conn.Err(); err != nil {
		return nil, err
	}

	return pool, nil
}
