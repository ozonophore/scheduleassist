//go:build !postgres

package adapter

import (
	"ScheduleAssist/internal/config"
	"ScheduleAssist/internal/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Initialize(cfg *config.Config) {
	logger.Info("SQLite adapter initialized")
	var err error
	db, err = gorm.Open(sqlite.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		logger.Panic("failed to connect to database: %v", err)
	}
	initPool()
	migration()
}

func GetDB() *gorm.DB {
	return db
}
