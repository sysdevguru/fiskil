package dcl

import (
	"context"
	"fmt"

	"github.com/sysdevguru/fiskil/models"
)

// Parse parses all logs provided from the worker and extracts severities
func (uc *UseCase) Parse(ctx context.Context, logs *[]models.Log) (*map[string]*models.Severity, error) {
	return GetSeverities(logs)
}

func GetSeverities(logs *[]models.Log) (*map[string]*models.Severity, error) {
	severityMap := make(map[string]*models.Severity)
	for _, log := range *logs {
		key := fmt.Sprintf("%s-%s", log.ServiceName, log.Severity)
		if severity, ok := severityMap[key]; ok {
			severity.Count++
			continue
		}

		s := &models.Severity{
			ServiceName: log.ServiceName,
			Severity:    log.Severity,
			Count:       1,
		}
		severityMap[key] = s
	}

	return &severityMap, nil
}
