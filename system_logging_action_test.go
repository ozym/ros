package ros

import (
	"testing"
)

func TestSystemLoggingAction(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{systemLoggingAction("name"), `/system logging action print detail where name="name"`},
		{setSystemLoggingAction("name", "key", "value"), `/system logging action set [find name="name"] key="value"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("system logging action client mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
