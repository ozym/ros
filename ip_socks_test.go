package ros

import (
	"testing"
)

func TestIpSocks(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{ipSocks(), `/ip socks print`},
		{setIpSocks("enabled", "yes"), `/ip socks set enabled="yes"`},
		{setIpSocks("enabled", "no"), `/ip socks set enabled="no"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("tool socks mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
