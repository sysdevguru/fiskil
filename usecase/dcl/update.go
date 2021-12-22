package dcl

import (
	"context"
	"log"

	"github.com/sysdevguru/fiskil/models"
)

// Update inserts logs and updates severities
func (uc *UseCase) Update(ctx context.Context, logs *[]models.Log, severities *[]models.Severity) error {
	log.Printf("storing %d logs, updating %d severities", len(*logs), len(*severities))
	err := uc.repo.UpdateLogs(ctx, logs, severities)
	if err != nil {
		log.Printf("error: %v", err)
	}
	return err
}
