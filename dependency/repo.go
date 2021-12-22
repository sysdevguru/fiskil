package dependency

import (
	"context"
	"reflect"

	"github.com/sysdevguru/fiskil/models"
	"gorm.io/gorm"
)

type GormRepo struct {
	db *gorm.DB
}

// UpdateLogs inserts logs and updates severities
func (g *GormRepo) UpdateLogs(ctx context.Context, logs *[]models.Log, severities *[]models.Severity) error {
	// start transaction
	tx := g.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			// rollback for the failure
			tx.Rollback()
		}
	}()

	err := tx.CreateInBatches(*logs, len(*logs)).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	newSeverities := []models.Severity{}

	for _, severity := range *severities {
		// get severity
		var dbSeverity models.Severity
		_ = tx.
			Where("service_name = ?", severity.ServiceName).
			Where("severity = ?", severity.Severity).
			Find(&dbSeverity)

		if reflect.DeepEqual(dbSeverity, severity) {
			newSeverities = append(newSeverities, severity)
			continue
		}

		severity.Count += dbSeverity.Count
		severity.ID = dbSeverity.ID
		newSeverities = append(newSeverities, severity)
	}

	err = tx.
		Where("service_name like 'service%'").
		Delete(models.Severity{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.CreateInBatches(newSeverities, len(newSeverities)).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// trigger transaction
	return tx.Commit().Error
}

func NewRepo(db *gorm.DB) *GormRepo {
	return &GormRepo{
		db: db,
	}
}
