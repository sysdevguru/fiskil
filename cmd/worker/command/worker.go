package command

import (
	"context"
	"log"

	"github.com/sysdevguru/fiskil/config"
	"github.com/sysdevguru/fiskil/task"
)

func Start(ctx context.Context, cfg config.Config) {
	// initialize worker pool
	service, err := task.NewService(cfg)
	if err != nil {
		log.Fatalf("unable to start worker %s", err.Error())
	}

	log.Println("starting application...")
	service.Start()
}
