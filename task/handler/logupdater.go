package handler

import (
	"log"
	"time"

	gocraft "github.com/gocraft/work"
	"github.com/sysdevguru/fiskil/models"
)

// ProcessNewLogUpdateTask updates database with provided logs/severities
func (c *Context) ProcessNewLogUpdateTask(job *gocraft.Job) error {
	log.Println("triggering log updater...")

	logsData := job.Args["logs"].([]interface{})
	logs := getLogs(logsData)

	severitiesData := job.Args["severities"].([]interface{})
	severities := getSeverities(severitiesData)
	return c.deps.logger.Update(c.ctx, logs, severities)
}

func getLogs(logsData []interface{}) *[]models.Log {
	logs := []models.Log{}

	for _, logData := range logsData {
		logMap := logData.(map[string]interface{})
		log := models.Log{}
		for k, v := range logMap {
			switch k {
			case "service_name":
				log.ServiceName = v.(string)
			case "severity":
				log.Severity = v.(string)
			case "payload":
				log.Payload = v.(string)
			case "timestamp":
				layout := "2006-01-02T15:04:05.000Z"
				timestamp := v.(string)
				t, err := time.Parse(layout, timestamp)
				if err != nil {
					continue
				}
				log.Timestamp = t
			default:
			}
		}
		logs = append(logs, log)
	}

	return &logs
}

func getSeverities(severitiesData []interface{}) *[]models.Severity {
	severities := []models.Severity{}

	for _, severityData := range severitiesData {
		severity := models.Severity{}
		severityMap := severityData.(map[string]interface{})
		for k, v := range severityMap {
			switch k {
			case "service_name":
				severity.ServiceName = v.(string)
			case "severity":
				severity.Severity = v.(string)
			case "count":
				severity.Count = int(v.(float64))
			case "created_at":
				layout := "2006-01-02T15:04:05.000Z"
				timestamp := v.(string)
				t, err := time.Parse(layout, timestamp)
				if err != nil {
					continue
				}
				severity.CreatedAt = t
			default:
			}
		}
		severities = append(severities, severity)
	}

	return &severities
}
