package adapter

import (
	"ScheduleAssist/internal/logger"
	"time"
)

func initPool() {
	sqlDB, err := GetDB().DB()
	if err != nil {
		logger.Panic("failed to get database connection: %v", err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(5)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(10)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	logger.Info("database connection pool initialized")
}
