package mapper

import (
	"ScheduleAssist/internal/model/domain"
	"ScheduleAssist/internal/model/orm"
)

func MapUserDBToUser(user *orm.User) *domain.User {
	return &domain.User{
		ID:         user.ID,
		Username:   user.Username,
		Password:   user.Password,
		TelegramID: user.TelegramID,
	}
}

func MapTaskToDB(task *domain.Task, userID uint32) *orm.Task {
	return &orm.Task{
		ID:                     task.ID,
		UserID:                 userID,
		CRON:                   task.CRON,
		TaskType:               orm.TaskType(task.TaskType),
		ShortTask:              task.ShortTask,
		FullTask:               task.FullTask,
		HumanReadableCron:      task.HumanReadableCron,
		CheckStatusCron:        task.CheckStatusCron,
		HumanReadableChackCron: task.HumanReadableChackCron,
		Amount:                 task.Amount,
		CreatedAt:              nil,
		UpdatedAt:              nil,
		StartDate:              task.StartDate.Time,
		EndDate:                task.EndDate.Time,
		Completed:              false,
		CompletionDate:         nil,
	}
}

func MapTasksToDB(tasks *[]domain.Task, userID uint32) *[]orm.Task {
	var result []orm.Task
	for _, task := range *tasks {
		value := MapTaskToDB(&task, userID)
		result = append(result, *value)
	}
	return &result
}
