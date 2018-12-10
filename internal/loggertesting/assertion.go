package loggertesting

import (
	"strings"
	"testing"

	"github.com/goph/logur"
)

const (
	// Makes the log event assertion skip raw line comparison
	SkipRawLine uint8 = 1 << iota

	// Matches lines without matching the last newline character
	AllowNoNewLine
)

// AssertLogEvents asserts that two LogEvents are identical.
func AssertLogEvents(t *testing.T, expected logur.LogEvent, actual logur.LogEvent, flags uint8) {
	t.Helper()

	if expected.Level != actual.Level {
		t.Errorf("expected log levels to be equal\ngot:  %s\nwant: %s", actual.Level, expected.Level)
	}

	if flags&AllowNoNewLine != 0 {
		if expected.Line != actual.Line && strings.TrimSuffix(expected.Line, "\n") != actual.Line {
			t.Errorf("expected log lines to be equal\ngot:  %q\nwant: %q", actual.Line, expected.Line)
		}
	} else {
		if expected.Line != actual.Line {
			t.Errorf("expected log lines to be equal\ngot:  %q\nwant: %q", actual.Line, expected.Line)
		}
	}

	if flags&SkipRawLine == 0 {
		if len(expected.RawLine) != len(actual.RawLine) {
			t.Errorf("expected raw log lines to be equal\ngot:  %v\nwant: %v", actual.RawLine, expected.RawLine)
		}

		for key, value := range expected.RawLine {
			if actual.RawLine[key] != value {
				t.Errorf("expected raw log lines to be equal\ngot:  %v\nwant: %v", actual.RawLine, expected.RawLine)

				break
			}
		}
	}

	if len(expected.Fields) != len(actual.Fields) {
		t.Errorf("expected log fields to be equal\ngot:  %v\nwant: %v", actual.Fields, expected.Fields)
	}

	for key, value := range expected.Fields {
		if actual.Fields[key] != value {
			t.Errorf("expected log fields to be equal\ngot:  %v\nwant: %v", actual.Fields, expected.Fields)

			break
		}
	}
}
