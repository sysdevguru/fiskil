package dependency

import (
	"github.com/sysdevguru/fiskil/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewGormWithPostgres initializes a Postgres database connection.
func NewGormWithPostgres(cfg config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.Database.URL), &gorm.Config{
		Logger: nil,
	})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConnections)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CloseDatabaseConnection cleans up the connection to the db.
func CloseDatabaseConnection(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
