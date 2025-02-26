package orm

import (
	"time"
)

type TaskType string

const (
	OneTime    TaskType = "one-time"
	Repeatable TaskType = "repeatable"
)

type User struct {
	ID         uint32  `gorm:"primaryKey;autoIncrement"`
	Username   *string `gorm:"size:255"`
	Password   *string `gorm:"size:255"`
	TelegramID *string `gorm:"size:255;uniqueIndex"`
}

// Task - основная модель задачи
type Task struct {
	ID                     uint32     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID                 uint32     `gorm:"not null"`
	User                   User       `gorm:"foreignKey:UserID;not null;"`
	TaskType               TaskType   `gorm:"task_type;size:15"`        // Тип задачи (one-time или repeatable)
	CRON                   string     `json:"cron"`                     // Расписание задачи в формате CRON
	ShortTask              string     `gorm:"unique" json:"short_task"` // Краткое описание задачи
	FullTask               string     `json:"full_task"`                // Полное описание задачи
	HumanReadableCron      string     `json:"human_readable_cron"`      // Человекочитаемое описание расписания CRON
	CheckStatusCron        string     `json:"check_status_cron"`        // Расписание проверки статуса задачи CRON
	HumanReadableChackCron string     `json:"human_readable_check_cron"`
	Amount                 int        `json:"amount"`          // Количество действий, если задано, для единичных операций = 1
	CreatedAt              *time.Time `gorm:"autoCreateTime"`  // Дата создания задачи
	UpdatedAt              *time.Time `gorm:"autoUpdateTime"`  // Дата последнего обновления задачи
	StartDate              time.Time  `json:"start_date"`      // Дата начала выполнения задачи
	EndDate                time.Time  `json:"end_date"`        // Дата окончания выполнения задачи
	Completed              bool       `json:"completed"`       // Статус выполнения задачи
	CompletionDate         *time.Time `json:"completion_date"` // Дата выполнения задачи
}
