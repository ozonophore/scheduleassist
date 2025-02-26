package mapper

import (
	"ScheduleAssist/internal/model/domain"
	"ScheduleAssist/internal/model/orm"
	time2 "ScheduleAssist/internal/time"
	"testing"
	"time"
)

func TestMapTasksToDB(t *testing.T) {
	tasks := []domain.Task{
		{
			ID:                     1,
			TaskType:               "repeatable",
			CRON:                   "0 8,18 * * *",
			ShortTask:              "Поливать цветы",
			FullTask:               "Поливать цветы два раза в день в 8 утра и 6 вечера",
			HumanReadableCron:      "Проверить статус в 20:00",
			CheckStatusCron:        "0 20 * * *",
			HumanReadableChackCron: "Проверить статус в 20:00",
			Amount:                 2,
			StartDate:              time2.CustomTime{Time: time.Date(2025, 2, 24, 8, 0, 0, 0, time.UTC)},
			EndDate:                time2.CustomTime{Time: time.Date(2026, 2, 24, 18, 0, 0, 0, time.UTC)},
		},
		{
			ID:                     2,
			TaskType:               "one-time",
			CRON:                   "",
			ShortTask:              "Купить продукты",
			FullTask:               "Купить продукты на неделю",
			HumanReadableCron:      "",
			CheckStatusCron:        "",
			HumanReadableChackCron: "",
			Amount:                 1,
			StartDate:              time2.CustomTime{Time: time.Date(2025, 2, 24, 8, 0, 0, 0, time.UTC)},
			EndDate:                time2.CustomTime{Time: time.Date(2025, 2, 24, 18, 0, 0, 0, time.UTC)},
		},
	}

	expected := []orm.Task{
		{
			ID:                     1,
			TaskType:               orm.TaskType("repeatable"),
			CRON:                   "0 8,18 * * *",
			ShortTask:              "Поливать цветы",
			FullTask:               "Поливать цветы два раза в день в 8 утра и 6 вечера",
			HumanReadableCron:      "Проверить статус в 20:00",
			CheckStatusCron:        "0 20 * * *",
			HumanReadableChackCron: "Проверить статус в 20:00",
			Amount:                 2,
			StartDate:              time.Date(2025, 2, 24, 8, 0, 0, 0, time.UTC),
			EndDate:                time.Date(2026, 2, 24, 18, 0, 0, 0, time.UTC),
			Completed:              false,
			CreatedAt:              nil,
			UpdatedAt:              nil,
			CompletionDate:         nil,
		},
		{
			ID:                     2,
			TaskType:               orm.TaskType("one-time"),
			CRON:                   "",
			ShortTask:              "Купить продукты",
			FullTask:               "Купить продукты на неделю",
			HumanReadableCron:      "",
			CheckStatusCron:        "",
			HumanReadableChackCron: "",
			Amount:                 1,
			StartDate:              time.Date(2025, 2, 24, 8, 0, 0, 0, time.UTC),
			EndDate:                time.Date(2025, 2, 24, 18, 0, 0, 0, time.UTC),
			Completed:              false,
			CreatedAt:              nil,
			UpdatedAt:              nil,
			CompletionDate:         nil,
		},
	}

	actual := MapTasksToDB(&tasks)
	if len(*actual) != len(expected) {
		t.Fatalf("expected %d tasks but got %d", len(expected), len(*actual))
	}

	for i, task := range *actual {
		if task != expected[i] {
			t.Errorf("expected task %v but got %v", expected[i], task)
		}
	}
}
