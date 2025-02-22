package model

import "time"

type TaskType string

const (
	OneTime    TaskType = "one-time"
	Repeatable TaskType = "repeatable"
)

// Task - основная модель задачи
type Task struct {
	ID                string    `json:"id"`
	TaskType          TaskType  `json:"task_type"`           // Тип задачи (one-time или repeatable)
	CRON              string    `json:"cron"`                // Расписание задачи в формате CRON
	ShortTask         string    `json:"short_task"`          // Краткое описание задачи
	FullTask          string    `json:"full_task"`           // Полное описание задачи
	HumanReadableCron string    `json:"human_readable_cron"` // Человекочитаемое описание расписания CRON
	CheckStatusCron   string    `json:"check_status_cron"`   // Расписание проверки статуса задачи CRON
	Amount            int       `json:"amount"`              // Количество действий, если задано, для единичных операций = 1
	CreatedAt         time.Time `json:"created_at"`          // Дата создания задачи
	UpdatedAt         time.Time `json:"updated_at"`          // Дата последнего обновления задачи
	StartDate         time.Time `json:"start_date"`          // Дата начала выполнения задачи
	EndDate           time.Time `json:"end_date"`            // Дата окончания выполнения задачи
	Completed         bool      `json:"completed"`           // Статус выполнения задачи
	CompletionDate    time.Time `json:"completion_date"`     // Дата выполнения задачи
}
