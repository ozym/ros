package ros

import (
	"testing"
)

func TestSystemResource(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{systemResource(), `/system resource print`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("system resource mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
