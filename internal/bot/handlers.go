package bot

import (
	"ScheduleAssist/internal/adapter"
	context2 "ScheduleAssist/internal/context"
	"ScheduleAssist/internal/logger"
	"ScheduleAssist/internal/model/mapper"
	"ScheduleAssist/internal/textanalyzer"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	logger.Info("Received message from %s: %s", message.From.UserName, message.Text)
	user, err := adapter.SetUserWithUsername(message.From.UserName, message.From.ID)
	if err != nil {
		logger.Panic("Error setting user: %v", err)
	}
	ctx, ok := context2.GetContextPool().GetContext(message.Chat.ID, mapper.MapUserDBToUser(user))
	logger.Debug("Context is new: %v", !ok)
	if !ok {
		ctx.OnClose = func(key int64) {
			logger.Info("Context with key %d has been closed", key)
			SendMessage(bot, key, "Время вашей сессии истекло. Пожалуйста, начните новый запрос или уточните информацию, чтобы продолжить.")
		}
		handleStart(bot, message.Chat.ID)
		return
	}
	switch ctx.CurrOperation {
	case context2.AddTaskConfirm:
		handleEditTasks(bot, message)
	}

	switch message.Text {
	case "/start":
		//SendMessage(bot, message.Chat.ID, "Привет! Я твой Telegram-бот!")
		handleStart(bot, message.Chat.ID)
	case "/tasks":
		SendMessage(bot, message.Chat.ID, "Список задач:\n1. Покормить кота\n2. Погулять с собак)")
	case "/add_task":
		SendMessage(bot, message.Chat.ID, "Добавить задачу")
	case "/statistic":
		handleStatistics(bot, message.Chat.ID)
	case "/help":
		SendMessage(bot, message.Chat.ID, "Доступные команды:\n"+
			"/start    - 🚀 Начать\n"+
			"/tasks    - 📋 Список задач\n"+
			"/add_task - ➕ Добавить задачу\n"+
			"/statistic - 📊 Статистика\n"+
			"/help     - ❓ Помощь")
	default:
		handleDefaultMessage(bot, message)
	}
}

func HandleCallbackQuery(bot *tgbotapi.BotAPI, query *tgbotapi.CallbackQuery) {
	logger.Info("Received callback query from %s: %s", query.From.UserName, query.Data)
	//var responseText string
	messageID := query.Message.MessageID
	switch query.Data {
	//case "week":
	//	responseText = "📅 Статистика за **неделю**"
	//case "month":
	//	responseText = "📆 Статистика за **месяц**"
	//case "two_months":
	//	responseText = "🗓️ Статистика за **2 месяца**"
	//case "three_months":
	//	responseText = "📊 Статистика за **3 месяца**"
	case "edit_tasks":
		handleEditTasks(bot, query.Message)
	case "save_tasks":
		handleSaveTasks(bot, query.Message)
	case "add_task":
		handleDefaultMessage(bot, query.Message)
	case "tasks":
		updateSubmenu(bot, query.Message.Chat.ID, messageID, "список задач")
	case "back":
		updateMainMenu(bot, query.Message.Chat.ID, messageID)
	}

	bot.Request(tgbotapi.NewCallback(query.ID, ""))
	//// Ответ на нажатие кнопки
	//callback := tgbotapi.NewCallback(query.ID, responseText)
	//bot.Send(callback)
	//
	//// Сообщение с результатом
	//msg := tgbotapi.NewMessage(query.Message.Chat.ID, responseText)
	//bot.Send(msg)
}

func SendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	bot.Send(msg)
}

func handleStart(bot *tgbotapi.BotAPI, chatID int64) {
	sendMainMenu(bot, chatID)
}

func sendMainMenu(bot *tgbotapi.BotAPI, chatID int64) {
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📋 Список задач", "tasks"),
			tgbotapi.NewInlineKeyboardButtonData("➕ Добавить задачу", "add_task"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, "Выберите действие:")
	msg.ReplyMarkup = buttons
	_, ok := context2.GetContextPool().GetContext(chatID, nil)
	bot.Send(msg)
	logger.Debug("Context is new: %v", !ok)
}

func updateMainMenu(bot *tgbotapi.BotAPI, chatID int64, messageID int) {
	edit := tgbotapi.NewEditMessageText(chatID, messageID, "Выберите действие:")
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📋 Список задач", "tasks"),
			tgbotapi.NewInlineKeyboardButtonData("➕ Добавить задачу", "add_task"),
		),
	)
	edit.ReplyMarkup = &buttons
	bot.Send(edit)
}

func updateSubmenu(bot *tgbotapi.BotAPI, chatID int64, messageID int, section string) {
	msg := tgbotapi.NewEditMessageText(chatID, messageID, "Вы выбрали "+section)
	msg.ReplyMarkup = &tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{tgbotapi.NewInlineKeyboardButtonData("⬅️ Назад", "back")},
		},
	}
	bot.Send(msg)
}

func handleStatistics(bot *tgbotapi.BotAPI, chatID int64) {
	// Кнопки выбора периода
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📅 За неделю", "week"),
			tgbotapi.NewInlineKeyboardButtonData("📆 За месяц", "month"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🗓️ За 2 месяца", "two_months"),
			tgbotapi.NewInlineKeyboardButtonData("📊 За 3 месяца", "three_months"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, "Выберите период статистики:")
	msg.ReplyMarkup = buttons
	bot.Send(msg)
}

func handleEditTasks(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	chatID := message.Chat.ID
	ctx := context2.GetContextPool().GetContextValue(chatID)
	if ctx.CurrOperation != context2.AddTaskConfirm {
		SendMessage(bot, chatID, "Нет данных для редактирования")
		return
	}

}

func handleDefaultMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	chatID := message.Chat.ID
	ctx := context2.GetContextPool().GetContextValue(chatID).SetOperation(context2.AddTask)
	textanalyzer.Context(ctx)
	tasks, content := textanalyzer.PrepareModel(ctx, message.Text)
	if tasks == nil {
		msg := tgbotapi.NewMessage(chatID, content)
		bot.Send(msg)
	} else {
		ctx.SetOperation(context2.AddTaskConfirm)
		ctx.SetTasks(tasks)
		msg := tgbotapi.NewMessage(chatID, textanalyzer.ToHTML(tasks))
		msg.ParseMode = tgbotapi.ModeHTML
		buttons := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Сохранить", "save_tasks"),
			),
		)
		msg.ReplyMarkup = buttons
		_, err := bot.Send(msg)
		if err != nil {
			logger.Error("Error sending message: %s", err.Error())
		}
		logger.Debug("Task: %s", tasks)
	}
}

func handleSaveTasks(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	chatID := message.Chat.ID
	ctx := context2.GetContextPool().GetContextValue(chatID)
	if ctx.CurrOperation != context2.AddTaskConfirm {
		SendMessage(bot, chatID, "Нет данных для сохранения")
		return
	}
	tasks := ctx.GetTasks()
	err := adapter.CreatTasks(mapper.MapTasksToDB(tasks, ctx.GetUserID()))
	if err != nil {
		SendMessage(bot, chatID, "Ошибка сохранения задач")
		logger.Error("Error saving tasks: %s", err.Error())
		return
	}
	SendMessage(bot, chatID, "Задачи сохранены")
	ctx.SetOperation(context2.None)
}
