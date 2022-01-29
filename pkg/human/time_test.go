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

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("duration < 60min", func(t *testing.T) {
		got := Duration(lastHour)
		want := "within an hour"
		assertCorrectMessage(t, got, want)
	})
	t.Run("duration < 24h", func(t *testing.T) {
		got := Duration(lastDay)
		want := "within a day"
		assertCorrectMessage(t, got, want)
	})
	t.Run("duration > 24h && < 7d", func(t *testing.T) {
		got := Duration(lastWeek)
		want := "within a week"
		assertCorrectMessage(t, got, want)
	})
	t.Run("duration > 7d && < 30d", func(t *testing.T) {
		got := Duration(lastMonth)
		want := "within 2 weeks"
		assertCorrectMessage(t, got, want)
	})
	t.Run("duration > 30d && < 1a", func(t *testing.T) {
		got := Duration(lastQuarter)
		want := "within 6 weeks"
		assertCorrectMessage(t, got, want)
	})
	t.Run("duration > 365d && < 2a", func(t *testing.T) {
		got := Duration(lastYear)
		want := "within 31 weeks"
		assertCorrectMessage(t, got, want)
	})
	t.Run("duration > 1a", func(t *testing.T) {
		got := Duration(lastDecade)
		want := "within a year"
		assertCorrectMessage(t, got, want)
	})
	t.Run("duration > 2a", func(t *testing.T) {
		got := Duration(lastDecade)
		want := "within a year"
		assertCorrectMessage(t, got, want)
	})
}
