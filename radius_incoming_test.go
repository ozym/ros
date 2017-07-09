package ros

import (
	"testing"
)

func TestRadiusIncoming(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{radiusIncoming(), `/radius incoming print`},
		{setRadiusIncomingAccept(true), `/radius incoming set accept="yes"`},
		{setRadiusIncomingAccept(false), `/radius incoming set accept="no"`},
		{setRadiusIncomingPort("3799"), `/radius incoming set port="3799"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("radius incoming mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
