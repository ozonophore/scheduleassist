// contextpool_test.go
package context

import (
	"testing"
	"time"
)

func TestNewContextPool(t *testing.T) {
	timer := 1 * time.Second
	cp := NewContextPool(timer)
	if cp == nil {
		t.Fatal("Expected non-nil ContextPool")
	}
	if cp.timer != timer {
		t.Errorf("Expected timer %v, got %v", timer, cp.timer)
	}
	if cp.pool == nil {
		t.Fatal("Expected non-nil pool map")
	}
}

func TestGetContext(t *testing.T) {
	timer := 1 * time.Second
	cp := NewContextPool(timer)
	key := int64(0)
	ctx, _ := cp.GetContext(key)
	if ctx == nil {
		t.Fatal("Expected non-nil context")
	}
	if _, ok := cp.pool[key]; !ok {
		t.Errorf("Expected key %v to be in pool", key)
	}
}

func TestClose(t *testing.T) {
	timer := 1 * time.Second
	cp := NewContextPool(timer)
	key := int64(0)
	cp.GetContext(key)
	cp.Close()
	if len(cp.pool) != 0 {
		t.Errorf("Expected pool to be empty, got %d", len(cp.pool))
	}
}

func TestResetClose(t *testing.T) {
	timer := 2 * time.Second
	cp := NewContextPool(timer)
	key := int64(0)
	cp.GetContext(key)
	time.Sleep(1 * time.Second)
	cp.GetContext(key)
	time.Sleep(1 * time.Second)
	if len(cp.pool) == 0 {
		t.Errorf("Expected pool to be non-empty, got %d", len(cp.pool))
	}
	cp.Close()
	if len(cp.pool) != 0 {
		t.Errorf("Expected pool to be empty, got %d", len(cp.pool))
	}
}
