package ros

import (
	"testing"
)

func TestRos_ParseBool(t *testing.T) {

	tests := []struct {
		v string
		b bool
	}{
		{"yes", true},
		{"no", false},
		{"yesno", false},
		{"", false},
	}

	for _, x := range tests {
		if r := ParseBool(x.v); r != x.b {
			t.Errorf("system format bool mismatch %s: got '%v', expected '%v'", x.v, r, x.b)
		}
	}
}

func TestRos_FormaBool(t *testing.T) {

	tests := []struct {
		b bool
		v string
	}{
		{true, "yes"},
		{false, "no"},
	}

	for _, x := range tests {
		if r := FormatBool(x.b); r != x.v {
			t.Errorf("system format bool mismatch %v: got '%s', expected '%s'", x.b, r, x.v)
		}
	}
}
