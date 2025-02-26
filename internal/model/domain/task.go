package domain

import (
	time2 "ScheduleAssist/internal/time"
	"time"
)

type TaskType string

const (
	OneTime    TaskType = "one-time"
	Repeatable TaskType = "repeatable"
)

type User struct {
	ID         uint32
	Username   *string
	Password   *string
	TelegramID *string
}

// Task - основная модель задачи
type Task struct {
	ID                     uint32           `json:"id"`
	TaskType               TaskType         `json:"task_type"`           // Тип задачи (one-time или repeatable)
	CRON                   string           `json:"cron"`                // Расписание задачи в формате CRON
	ShortTask              string           `json:"short_task"`          // Краткое описание задачи
	FullTask               string           `json:"full_task"`           // Полное описание задачи
	HumanReadableCron      string           `json:"human_readable_cron"` // Человекочитаемое описание расписания CRON
	CheckStatusCron        string           `json:"check_status_cron"`   // Расписание проверки статуса задачи CRON
	HumanReadableChackCron string           `json:"human_readable_check_cron"`
	Amount                 int              `json:"amount"`     // Количество действий, если задано, для единичных операций = 1
	UpdatedAt              time.Time        `json:"updated_at"` // Дата последнего обновления задачи
	StartDate              time2.CustomTime `json:"start_date"` // Дата начала выполнения задачи
	EndDate                time2.CustomTime `json:"end_date"`   // Дата окончания выполнения задачи
}
