package context

import (
	"ScheduleAssist/internal/model/domain"
	"context"
	"github.com/sashabaranov/go-openai"
	"sync"
	"time"
)

type OperationType string

const (
	None           OperationType = "none"
	AddTask        OperationType = "add_task"
	AddTaskConfirm OperationType = "add_task_confirm"
	ShowTasks      OperationType = "show_tasks"
)

type AutoCancelContext struct {
	sync.Mutex
	ctx           context.Context
	cancel        context.CancelFunc
	duration      *time.Duration
	timer         *time.Timer
	chatID        int64
	OnClose       func(chatId int64)
	CurrOperation OperationType
	request       *openai.ChatCompletionRequest
	tasks         *[]domain.Task
	user          *domain.User
}

func NewAutoCancelContext(duration time.Duration, chatID int64, user *domain.User) *AutoCancelContext {
	ctx, cancel := context.WithCancel(context.Background())
	return &AutoCancelContext{
		ctx:           ctx,
		cancel:        cancel,
		duration:      &duration,
		chatID:        chatID,
		CurrOperation: None,
		tasks:         nil,
		user:          user,
		timer: time.AfterFunc(duration, func() {
			cancel()
		}),
	}
}

func (acc *AutoCancelContext) GetUserID() uint32 {
	return acc.user.ID
}

func (acc *AutoCancelContext) Reset() {
	acc.Lock()
	defer acc.Unlock()
	acc.timer.Reset(*acc.duration)
}

func (acc *AutoCancelContext) SetOperation(operation OperationType) *AutoCancelContext {
	acc.CurrOperation = operation
	if operation == None {
		acc.SetRequest(nil)
		acc.tasks = nil
	}
	return acc
}

func (acc *AutoCancelContext) GetTasks() *[]domain.Task {
	return acc.tasks
}

func (acc *AutoCancelContext) SetTasks(tasks *[]domain.Task) {
	acc.tasks = tasks
}

func (acc *AutoCancelContext) GetRequest() *openai.ChatCompletionRequest {
	return acc.request
}

func (acc *AutoCancelContext) SetRequest(request *openai.ChatCompletionRequest) {
	acc.request = request
}

func (acc *AutoCancelContext) GetContext() context.Context {
	return acc.ctx
}
