package human

import (
	"testing"
	"time"
)

func TestHello(t *testing.T) {
	now := time.Now()

	lastHour := time.Since(now.Add(-10 * time.Minute))
	lastDay := time.Since(now.Add(-10 * time.Hour))
	lastWeek := time.Since(now.AddDate(0, 0, -3))
	lastMonth := time.Since(now.AddDate(0, 0, -10))
	lastQuarter := time.Since(now.AddDate(0, -1, -10))
	lastYear := time.Since(now.AddDate(0, -7, 0))
	lastDecade := time.Since(now.AddDate(-1, -1, 0))

	timeTests := []struct {
		timestamp time.Duration
		want      string
	}{
		{lastHour, "within an hour"},
		{lastDay, "within a day"},
		{lastWeek, "within a week"},
		{lastMonth, "within 2 weeks"},
		{lastQuarter, "within 6 weeks"},
		{lastYear, "within 31 weeks"},
		{lastDecade, "within a year"},
		{lastDecade, "within a year"},
	}

	for _, tt := range timeTests {
		got := Duration(tt.timestamp)
		if got != tt.want {
			t.Errorf("got %q want %q", got, tt.want)
		}
	}

}
