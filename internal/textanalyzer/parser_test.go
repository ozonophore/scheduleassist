package textanalyzer

import "testing"

func TestExtractJSONFromText(t *testing.T) {
	text := "Спасибо за уточнение. Я создам задачу на полив цветов дважды в день с 8:00 до 6:00 в течение года. Расписание проверки статуса задачи будет установлено на 20:00 каждый день.\n\nВот сформированная задача:\n" +
		"\n```json\n{\n  \"task_type\": \"repeatable\",\n  \"short_task\": \"Поливать цветы\",\n  \"full_task\": \"Поливать цветы два раза в день в 8 утра и 6 вечера\",\n  \"amount\": 2,\n  \"cron\": \"0 8,18 * * *\",\n  \"human_readable_cron\": \"Проверить статус в 20:00\",\n  \"check_status_cron\": \"0 20 * * *\",\n  \"start_date\": \"2025-02-24T08:00:00Z\",\n  \"end_date\": \"2026-02-24T18:00:00Z\",\n  \"completed\": false\n}\n```"
	expected := "[{\n  \"task_type\": \"repeatable\",\n  \"short_task\": \"Поливать цветы\",\n  \"full_task\": \"Поливать цветы два раза в день в 8 утра и 6 вечера\",\n  \"amount\": 2,\n  \"cron\": \"0 8,18 * * *\",\n  \"human_readable_cron\": \"Проверить статус в 20:00\",\n  \"check_status_cron\": \"0 20 * * *\",\n  \"start_date\": \"2025-02-24T08:00:00Z\",\n  \"end_date\": \"2026-02-24T18:00:00Z\",\n  \"completed\": false\n}]"
	actual := ExtractJSONFromText(text)
	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}

func TestExtractJSONFromText_MultipleObjects(t *testing.T) {
	text := "Вот задачи:\n\n```json\n[\n  {\n    \"task_type\": \"repeatable\",\n    \"short_task\": \"Поливать цветы\",\n    \"full_task\": \"Поливать цветы два раза в день в 8 утра и 6 вечера\",\n    \"amount\": 2,\n    \"cron\": \"0 8,18 * * *\",\n    \"human_readable_cron\": \"Проверить статус в 20:00\",\n    \"check_status_cron\": \"0 20 * * *\",\n    \"start_date\": \"2025-02-24T08:00:00Z\",\n    \"end_date\": \"2026-02-24T18:00:00Z\",\n    \"completed\": false\n  },\n  {\n    \"task_type\": \"one-time\",\n    \"short_task\": \"Купить продукты\",\n    \"full_task\": \"Купить продукты на неделю\",\n    \"amount\": 1,\n    \"cron\": \"\",\n    \"human_readable_cron\": \"\",\n    \"check_status_cron\": \"\",\n    \"start_date\": \"2025-02-24T08:00:00Z\",\n    \"end_date\": \"2025-02-24T18:00:00Z\",\n    \"completed\": false\n  }\n]\n```"
	expected := "[\n  {\n    \"task_type\": \"repeatable\",\n    \"short_task\": \"Поливать цветы\",\n    \"full_task\": \"Поливать цветы два раза в день в 8 утра и 6 вечера\",\n    \"amount\": 2,\n    \"cron\": \"0 8,18 * * *\",\n    \"human_readable_cron\": \"Проверить статус в 20:00\",\n    \"check_status_cron\": \"0 20 * * *\",\n    \"start_date\": \"2025-02-24T08:00:00Z\",\n    \"end_date\": \"2026-02-24T18:00:00Z\",\n    \"completed\": false\n  },\n  {\n    \"task_type\": \"one-time\",\n    \"short_task\": \"Купить продукты\",\n    \"full_task\": \"Купить продукты на неделю\",\n    \"amount\": 1,\n    \"cron\": \"\",\n    \"human_readable_cron\": \"\",\n    \"check_status_cron\": \"\",\n    \"start_date\": \"2025-02-24T08:00:00Z\",\n    \"end_date\": \"2025-02-24T18:00:00Z\",\n    \"completed\": false\n  }\n]"
	actual := ExtractJSONFromText(text)
	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}

func TestExtractJSONFromText2(t *testing.T) {
	text := "```json\n[{\n  \"task_type\": \"repeatable\",\n  \"short_task\": \"Читать книгу\",\n  \"full_task\": \"Читать книгу каждый день 30 минут в течении дня\",\n  \"amount\": 1,\n  \"cron\": \"0 */30 * * *\",\n  \"human_readable_cron\": \"Каждый день каждые 30 минут\",\n  \"check_status_cron\": \"0 20 * * *\",\n  \"human_readable_check_cron\": \"Проверить статус в 20:00\",\n  \"start_date\": \"2025-02-24T00:00:00Z\",\n  \"end_date\": \"2026-02-24T00:00:00Z\",\n  \"completed\": false\n}]\n```"
	actual := ExtractJSONFromText(text)
	expected := "[{\n  \"task_type\": \"repeatable\",\n  \"short_task\": \"Читать книгу\",\n  \"full_task\": \"Читать книгу каждый день 30 минут в течении дня\",\n  \"amount\": 1,\n  \"cron\": \"0 */30 * * *\",\n  \"human_readable_cron\": \"Каждый день каждые 30 минут\",\n  \"check_status_cron\": \"0 20 * * *\",\n  \"human_readable_check_cron\": \"Проверить статус в 20:00\",\n  \"start_date\": \"2025-02-24T00:00:00Z\",\n  \"end_date\": \"2026-02-24T00:00:00Z\",\n  \"completed\": false\n}]"
	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}

func TestExtractJSONFromText3(t *testing.T) {
	text := "[{\n  \"task_type\": \"repeatable\",\n  \"short_task\": \"Читать книгу\",\n  \"full_task\": \"Читать книгу каждый день по 30 минут в течение дня\",\n  \"amount\": 365,\n  \"cron\": \"0 0,12 * * *\",\n  \"human_readable_cron\": \"Каждый день в полночь и в полдень\",\n  \"human_readable_check_cron\": \"Проверить статус в 20:00\",\n  \"check_status_cron\": \"0 20 * * *\",\n  \"start_date\": \"2025-02-24\",\n  \"end_date\": \"2026-02-24\",\n  \"completed\": false\n}]"
	actual := ExtractJSONFromText(text)
	println(actual)
}
