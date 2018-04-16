package ros

import (
	"testing"
)

func TestIpNeighbor(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{ipNeighbors(), `/ip neighbor print detail`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("ip neighbor mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
