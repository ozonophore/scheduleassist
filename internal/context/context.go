package context

import (
	"context"
	"sync"
	"time"
)

type AutoCancelContext struct {
	sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	duration *time.Duration
	timer    *time.Timer
}

func NewAutoCancelContext(duration time.Duration) *AutoCancelContext {
	ctx, cancel := context.WithCancel(context.Background())
	return &AutoCancelContext{
		ctx:      ctx,
		cancel:   cancel,
		duration: &duration,
		timer: time.AfterFunc(duration, func() {
			cancel()
		}),
	}
}

func (acc *AutoCancelContext) Reset() {
	acc.Lock()
	defer acc.Unlock()
	acc.timer.Reset(*acc.duration)
}
