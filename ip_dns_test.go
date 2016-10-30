package ros

import (
	"testing"
)

func TestIPDNS(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{ipDNS(), `/ip dns print`},
		{setIPDNS("key", "value"), `/ip dns set key="value"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("ip dns mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
