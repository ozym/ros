package ros

import (
	"testing"
)

func TestToolMacServerPing(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{toolMacServerPing(), `/tool mac-server ping print`},
		{setToolMacServerPing("enabled", "yes"), `/tool mac-server ping set enabled="yes"`},
		{setToolMacServerPing("enabled", "no"), `/tool mac-server ping set enabled="no"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("tool mac-server ping mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
