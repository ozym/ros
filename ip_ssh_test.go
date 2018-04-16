package ros

import (
	"testing"
)

func TestIpSsh(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{ipSsh(), `/ip ssh print`},
		{setIpSsh("key", "value"), `/ip ssh set key="value"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("ip ssh mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
