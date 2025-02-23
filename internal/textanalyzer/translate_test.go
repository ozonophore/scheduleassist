package textanalyzer

import (
	"ScheduleAssist/internal/model"
	"testing"
	"time"
)

func TestToMD(t *testing.T) {
	tasks := []model.Task{
		{
			TaskType:               "one-time",
			ShortTask:              "Покормить кота",
			FullTask:               "Покормить кота в 8 утра",
			Amount:                 1,
			CRON:                   "0 8 * * *",
			HumanReadableCron:      "Выполнить в 8:00",
			HumanReadableChackCron: "Проверить статус в 20:00",
			CheckStatusCron:        "0 20 * * *",
			StartDate:              time.Date(2025, 2, 24, 8, 0, 0, 0, time.UTC),
			EndDate:                time.Date(2025, 2, 24, 8, 0, 0, 0, time.UTC),
			Completed:              false,
		},
	}
	expected := "### Задача: Покормить кота\n" +
		"**Тип задачи**: одноразовая\n" +
		"Полное описание: Покормить кота в 8 утра\n" +
		"**Количество повторений**: 1\n" +
		"Расписание: Выполнить в 8:00\n" +
		"Проверит статус: Проверить статус в 20:00\n" +
		"Дата начала: 2025-02-24\n" +
		"Дата окончания: 2025-02-24\n" +
		"Завершено: false\n" +
		"\n"
	actual := ToHTML(&tasks)
	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}
