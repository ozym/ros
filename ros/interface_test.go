package ros

import (
	"testing"
)

func TestInterface(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{interfaces(), `/interface print detail`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("interface mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
