package dcl

import (
	"context"
	"log"
)

func (uc *UseCase) Fail(ctx context.Context, stage string) {
	log.Fatalf("failed job: %s", stage)
}
