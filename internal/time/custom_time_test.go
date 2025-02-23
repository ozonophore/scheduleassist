package time

import (
	"ScheduleAssist/internal/logger"
	"testing"
	"time"
)

func TestCustomTime_UnmarshalJSON(t *testing.T) {
	logger.Initialize(true)
	value := "2026-01-02T15:04:05Z07:00"
	ct := CustomTime{}
	err := ct.UnmarshalJSON([]byte(value))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expected := "2026-01-02"
	if ct.Format(time.DateOnly) != expected {
		t.Errorf("expected %s but got %s", expected, ct.Format(time.DateOnly))
	}
}
