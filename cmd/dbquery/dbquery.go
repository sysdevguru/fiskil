package main

import (
	"context"
	"log"
	"time"

	"github.com/sysdevguru/fiskil/config"
	"github.com/sysdevguru/fiskil/dependency"
	"github.com/sysdevguru/fiskil/models"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to read configuration: %v", err)
	}
	db, err := dependency.NewGormWithPostgres(cfg)
	if err != nil {
		log.Fatalf("failed to create database connection: %v", err)
	}

	ctx := context.Background()
	for {
		// get logs from the db
		logs := getLogs(ctx, db)
		log.Printf("found %d logs in db\n", len(*logs))

		// get severities
		severities := getSeverities(ctx, db)
		log.Printf("found %d severities in db\n", len(*severities))
		for _, severity := range *severities {
			log.Printf("%s : %s => %d", severity.ServiceName, severity.Severity, severity.Count)
		}
		log.Println()

		time.Sleep(30 * time.Second)
	}
}

func getLogs(ctx context.Context, db *gorm.DB) *[]models.Log {
	var logs []models.Log

	tx := db.WithContext(ctx)
	result := tx.
		Find(&logs)

	if result.Error != nil {
		return nil
	}

	return &logs
}

func getSeverities(ctx context.Context, db *gorm.DB) *[]models.Severity {
	var severities []models.Severity

	tx := db.WithContext(ctx)
	result := tx.
		Find(&severities)

	if result.Error != nil {
		return nil
	}

	return &severities
}
