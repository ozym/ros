package ros

import (
	"testing"
)

func TestIpArp(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{ipArps(), `/ip arp print detail`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("ip arp mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
