package ros

import (
	"testing"
)

func TestToolBandwidthServer(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{toolBandwidthServer(), `/tool bandwidth-server print`},
		{setToolBandwidthServer("disabled", "yes"), `/tool bandwidth-server set disabled="yes"`},
		{setToolBandwidthServer("disabled", "no"), `/tool bandwidth-server set disabled="no"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("tool bandwidth-server mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
