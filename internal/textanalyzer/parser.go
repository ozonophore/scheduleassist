package textanalyzer

import (
	"ScheduleAssist/internal/config"
	"ScheduleAssist/internal/logger"
	"ScheduleAssist/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"regexp"
	"strings"
	"time"
)

type IContext interface {
	GetRequest() *openai.ChatCompletionRequest
	SetRequest(request *openai.ChatCompletionRequest)
	GetContext() context.Context
}

var client *openai.Client

func Initialize(cfg *config.Config) {
	client = openai.NewClient(cfg.OpenAIToken)
}

func CreateRequestWithRoleSystem() *openai.ChatCompletionRequest {
	return &openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: fmt.Sprintf("Сейчас %s", time.Now().Format("2006-01-02 15:04:05")),
			},
			{
				Role: openai.ChatMessageRoleSystem,
				Content: "Ты эксперт по пормированию задач, разобей все задачи на следующие части: " +
					"тип задачи (one-time или repeatable), " +
					"краткое описание, " +
					"полное описание, " +
					"расписание в формате CRON(если не задано время выполнения уточнить), " +
					"дата начала, " +
					"дата окончания. " +
					"Расписание проверки статуса задачи CRON. " +
					"Если какое-то поле не заполнено, уточни у пользователя. " +
					"Если задача одноразовая дата начала ровна дате окончания. " +
					"Задача должна включать следующие поля: 'task_type' (one-time или repeatable), " +
					"'short_task' (краткое описание задачи), " +
					"'full_task' (полное описание задачи), " +
					"'amount' (Количество действий, если задано, для единичных операций = 1), " +
					"'cron' (формат расписания задачи), " +
					"'human_readable_cron' (человекочитаемый формат расписания выполнения задача), " +
					"'human_readable_check_cron' (человекочитаемый формат расписания проверки статуса задача), " +
					"'check_status_cron' (формат расписания проверки статуса задачи), " +
					"'start_date' (дата начала в формате YYYY-MM-DD), " +
					"'end_date' (дата окончания), " +
					"'completed' (статус выполнения в формате YYYY-MM-DD). " +
					" Если это несколько задач, то уточняй данные по все сразу. " +
					" Если точное время выполнения задачи не указано, указыват только день. " +
					"Если какая-то информация отсутствует, уточни у пользователя. " +
					"Если расписание проверки задачи не задано, то проверять в день выполенния задачи в 8 часов вечера. " +
					"Ответ верни в только в формате JSON без лишних деталей, точно как в примере. Коментарии в JSON запрещены. " +
					"Для фразы 'Покормить кода в 8 утра " +
					"Пример: \n" +
					"[{\n" +
					"  \"task_type\": \"one-time\",\n" +
					"  \"short_task\": \"Покормить кота\",\n" +
					"  \"full_task\": \"Покормить кота в 8 утра\",\n" +
					"  \"amount\": 1,\n" +
					"  \"cron\": \"0 8 * * *\",\n" +
					"  \"human_readable_cron\": \"Выполнить в 8:00\", " +
					"  \"check_status_cron\": \"0 20 * * *\",\n" +
					"  \"human_readable_check_cron\": \"Проверить статус в 20:00\", " +
					"  \"start_date\": \"2022-01-01\",\n" +
					"  \"end_date\": \"2022-01-01\",\n" +
					"  \"completed\": false\n" +
					"}]",
			},
		},
	}
}

func Context(parent IContext) {
	if parent.GetRequest() == nil {
		parent.SetRequest(CreateRequestWithRoleSystem())
	}
}

func PrepareModel(ctx IContext, text string) (taskPoint *[]model.Task, question string) {
	request := ctx.GetRequest()
	if request == nil {
		logger.Error("Request doesn't serialize")
	}
	message := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: text,
	}
	request.Messages = append(request.Messages, message)
	resp, err := client.CreateChatCompletion(ctx.GetContext(), *request)
	content := resp.Choices[0].Message.Content
	if err != nil {
		log.Fatal(err)
	}
	var tasks []model.Task
	content = ExtractJSONFromText(content)
	err = json.Unmarshal([]byte(content), &tasks)
	if err != nil {
		logger.Debug("%v", err)
		question = content
	} else {
		logger.Debug("Task: %s", content)
		taskPoint = &tasks
	}
	request.Messages = append(request.Messages, resp.Choices[0].Message)
	return
}

func ExtractJSONFromText(text string) string {
	// Регулярка для поиска JSON массива
	re := regexp.MustCompile(`(?s)(\{.*\}|\[.*\])`)
	// Находим совпадение
	matches := re.FindString(text)
	if matches == "" {
		return text
	}

	if !(strings.HasPrefix(matches, "[") && strings.HasSuffix(matches, "]")) {
		return fmt.Sprintf("[%s]", matches)
	}
	return matches
}
