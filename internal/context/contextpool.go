package context

import (
	"context"
	"sync"
	"time"
)

var instance *ContextPool

type ContextPool struct {
	sync.RWMutex
	timer time.Duration
	pool  map[int64]*AutoCancelContext
}

func NewContextPool(timer time.Duration) *ContextPool {
	pool := &ContextPool{
		pool:  make(map[int64]*AutoCancelContext),
		timer: timer,
	}
	instance = pool
	go pool.watchdog()
	return pool
}

func GetContextPool() *ContextPool {
	return instance
}

func (cp *ContextPool) GetContext(key int64) (context.Context, bool) {
	cp.Lock()
	defer cp.Unlock()
	if actx, ok := cp.pool[key]; ok {
		actx.Reset()
		return actx.ctx, true
	}
	actx := NewAutoCancelContext(cp.timer)
	cp.pool[key] = actx
	return actx.ctx, false
}

func (cp *ContextPool) watchdog() {
	for {
		for key, actx := range cp.pool {
			select {
			case <-actx.ctx.Done():
				func() {
					cp.Lock()
					defer cp.Unlock()
					delete(cp.pool, key)
				}()
			}
		}
		time.Sleep(cp.timer)
	}
}

func (cp *ContextPool) Close() {
	for _, actx := range cp.pool {
		actx.cancel()
	}
	// Wait for the pool to become empty
	for {
		cp.Lock()
		if len(cp.pool) == 0 {
			cp.Unlock()
			break
		}
		cp.Unlock()
		time.Sleep(10 * time.Millisecond)
	}
}
