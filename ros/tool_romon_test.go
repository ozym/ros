package ros

import (
	"testing"
)

func TestSystemRomon(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{toolRomon(true), `/romon print`},
		{toolRomon(false), `/tool romon print`},
		{setToolRomon("key", "value", true), `/romon set key="value"`},
		{setToolRomon("key", "value", false), `/tool romon set key="value"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("tool romon mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
