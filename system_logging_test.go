package ros

import (
	"testing"
)

func TestSystemLogging(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{systemLogging("action", "topics"), `/system logging print detail where action="action" topics="topics"`},
		{addSystemLogging("action", "topics"), `:if (:len [/system logging find action="action" topics="topics"] = 0) do={/system logging add action="action" topics="topics"}`},
		{removeSystemLogging("action", "topics"), `/system logging remove [find action="action" topics="topics"]`},
		{setSystemLoggingPrefix("action", "topics", "prefix"), `/system logging set [find action="action" topics="topics"] prefix="prefix"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("system logging mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
