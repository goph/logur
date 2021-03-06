package logtesting

import (
	"testing"

	"github.com/goph/logur"
)

// AssertLogEventsEqual asserts that two LogEvents are identical.
func AssertLogEventsEqual(t *testing.T, expected logur.LogEvent, actual logur.LogEvent) {
	t.Helper()

	err := logur.LogEventsEqual(expected, actual)
	if err != nil {
		t.Error(err)
	}
}
