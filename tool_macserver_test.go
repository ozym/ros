package ros

import (
	"testing"
)

func TestToolMacServer(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{toolMacServer(true), `/tool mac-server print detail where default`},
		{toolMacServer(false), `/tool mac-server print detail`},
		{setToolMacServer("disabled", "yes", true), `/tool mac-server set [find default] disabled="yes"`},
		{setToolMacServer("disabled", "no", true), `/tool mac-server set [find default] disabled="no"`},
		{setToolMacServer("disabled", "yes", false), `/tool mac-server set disabled="yes"`},
		{setToolMacServer("disabled", "no", false), `/tool mac-server set disabled="no"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("tool mac-server mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
