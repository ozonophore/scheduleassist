package main

import (
	"ScheduleAssist/internal/bot"
	"ScheduleAssist/internal/config"
	"ScheduleAssist/internal/context"
	"ScheduleAssist/internal/logger"
	"ScheduleAssist/internal/textanalyzer"
	"time"
)

func main() {
	cfg := config.InitConfig()
	logger.Initialize(cfg.Debug)
	textanalyzer.Initialize(cfg)
	context.NewContextPool(30 * time.Second)
	bot.StartBot(cfg)
}
