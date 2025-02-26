package textanalyzer

import (
	"ScheduleAssist/internal/model/domain"
	time2 "ScheduleAssist/internal/time"
	"testing"
	"time"
)

func TestToMD(t *testing.T) {
	tasks := []domain.Task{
		{
			TaskType:               "one-time",
			ShortTask:              "Покормить кота",
			FullTask:               "Покормить кота в 8 утра",
			Amount:                 1,
			CRON:                   "0 8 * * *",
			HumanReadableCron:      "Выполнить в 8:00",
			HumanReadableChackCron: "Проверить статус в 20:00",
			CheckStatusCron:        "0 20 * * *",
			StartDate:              time2.NewCustomTime(time.Date(2025, 2, 24, 8, 0, 0, 0, time.UTC)),
			EndDate:                time2.NewCustomTime(time.Date(2025, 2, 24, 8, 0, 0, 0, time.UTC)),
			Completed:              false,
		},
	}
	expected := "<b>Задача: Покормить кота</b>\n" +
		"<b>Тип задачи</b>: одноразовая\n" +
		"Полное описание: Покормить кота в 8 утра\n" +
		"<b>Количество повторений</b>: 1\n" +
		"Расписание: Выполнить в 8:00\n" +
		"Проверит статус: Проверить статус в 20:00\n" +
		"Дата начала: 2025-02-24\n" +
		"Дата окончания: 2025-02-24\n" +
		"\n"
	actual := ToHTML(&tasks)
	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}
