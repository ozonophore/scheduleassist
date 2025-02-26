//go:build postgres

package adapter

import (
	"ScheduleAssist/internal/config"
	"ScheduleAssist/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Initialize(cfg *config.Config) {
	logger.Info("Postgres adapter initialized")
	var err error
	db, err = gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		logger.Panic("failed to connect to database: %v", err)
	}
	initPool()
	migration()
}

func GetDB() {
	return db
}
