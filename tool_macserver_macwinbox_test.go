package ros

import (
	"testing"
)

func TestToolMacServerMacWinbox(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{toolMacServerMacWinbox(true), `/tool mac-server mac-winbox print detail where default`},
		{toolMacServerMacWinbox(false), `/tool mac-server mac-winbox print detail`},
		{setToolMacServerMacWinbox("disabled", "yes", true), `/tool mac-server mac-winbox set [find default] disabled="yes"`},
		{setToolMacServerMacWinbox("disabled", "no", true), `/tool mac-server mac-winbox set [find default] disabled="no"`},
		{setToolMacServerMacWinbox("disabled", "yes", false), `/tool mac-server mac-winbox set disabled="yes"`},
		{setToolMacServerMacWinbox("disabled", "no", false), `/tool mac-server mac-winbox set disabled="no"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("tool mac-server mac-winbox mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
