package bot

import (
	"ScheduleAssist/internal/config"
	"ScheduleAssist/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(cfg *config.Config) {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		logger.Panic(err)
	}

	bot.Debug = cfg.Debug
	logger.Info("Authorized on account %s", bot.Self.UserName)

	SetBotCommands(bot)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			go HandleMessage(bot, update.Message)
		}
		if update.CallbackQuery != nil {
			go HandleCallbackQuery(bot, update.CallbackQuery)
		}
	}
}

func SetBotCommands(bot *tgbotapi.BotAPI) {
	commands := []tgbotapi.BotCommand{
		{Command: "start", Description: "🚀 Запустить бота"},
		{Command: "tasks", Description: "📋 Список задач"},
		{Command: "add_task", Description: "➕ Добавить задачу"},
		{Command: "statistic", Description: "📊 Статистика"},
		{Command: "help", Description: "❓ Помощь"},
	}

	cmdCfg := tgbotapi.NewSetMyCommands(commands...)
	if _, err := bot.Request(cmdCfg); err != nil {
		logger.Error("Ошибка установки команд: %s", err.Error())
	}
}
