package time

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

const (
	dateFormat     = "2006-01-02"
	dateTimeFormat = "2006-01-02T15:04:05Z07:00"
)

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = strings.TrimSuffix(strings.TrimPrefix(str, "\""), "\"")

	re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}`)
	match := re.FindString(str)
	if match != "" {
		tm, err := time.Parse(time.DateOnly, match)
		if err != nil {
			return fmt.Errorf("неверный формат даты: %s", str)
		}
		ct.Time = tm
		return nil
	}
	return fmt.Errorf("неверный формат даты: %s", str)
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%s", ct.Format(time.DateOnly))), nil
}
