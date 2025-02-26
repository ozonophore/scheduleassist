package textanalyzer

import (
	"ScheduleAssist/internal/model/domain"
	"fmt"
	"strings"
)

func convertTaskType(taskType domain.TaskType) string {
	if taskType == "one-time" {
		return "одноразовая"
	}
	return "повторяющаяся"
}

func ToHTML(tasks *[]domain.Task) string {
	var sb strings.Builder
	for _, task := range *tasks {
		sb.WriteString(fmt.Sprintf("<b>Задача: %s</b>\n", task.ShortTask))
		sb.WriteString(fmt.Sprintf("<b>Тип задачи</b>: %s\n", convertTaskType(task.TaskType)))
		sb.WriteString(fmt.Sprintf("Полное описание: %s\n", task.FullTask))
		sb.WriteString(fmt.Sprintf("<b>Количество повторений</b>: %d\n", task.Amount))
		sb.WriteString(fmt.Sprintf("Расписание: %s\n", task.HumanReadableCron))
		sb.WriteString(fmt.Sprintf("Проверит статус: %s\n", task.HumanReadableChackCron))
		sb.WriteString(fmt.Sprintf("Дата начала: %s\n", task.StartDate.Format("2006-01-02")))
		sb.WriteString(fmt.Sprintf("Дата окончания: %s\n", task.EndDate.Format("2006-01-02")))
		sb.WriteString("\n")
	}
	return sb.String()
}
