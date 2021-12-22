package handler

import (
	"encoding/json"
	"log"

	gocraft "github.com/gocraft/work"
	"github.com/sysdevguru/fiskil/config"
	"github.com/sysdevguru/fiskil/models"
)

// ProcessNewLogParseTask triggers log pulling and parsing work
func (c *Context) ProcessNewLogParseTask(job *gocraft.Job) error {
	log.Println("triggering log parser...")
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	logs := []models.Log{}

	// pull logs from the pubsub
	batch := c.conn.ReadBatch(cfg.Kafka.ReadMin, cfg.Kafka.ReadMax)
	b := make([]byte, cfg.Kafka.ReadMin)
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}

		log := models.Log{}
		if err := json.Unmarshal(b[:n], &log); err != nil {
			return err
		}
		logs = append(logs, log)

		// update with defined amount
		if len(logs) >= cfg.App.BatchRecordCount {
			break
		}
	}

	log.Printf("read %d logs", len(logs))

	if len(logs) == 0 {
		return nil
	}

	severityMap, err := c.deps.logger.Parse(c.ctx, &logs)
	if err != nil {
		return err
	}

	severities := []models.Severity{}
	for _, severity := range *severityMap {
		severities = append(severities, *severity)
	}

	// trigger next task, log updater
	err = c.deps.enqueuer.AddNewLogUpdateTask(&logs, &severities)
	return err
}
