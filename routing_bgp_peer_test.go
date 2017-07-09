package ros

import (
	"testing"
)

func TestRoutingBgpPeer(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{routingBgpPeer("address"), `/routing bgp peer print detail where remote-address="address"`},
		{setRoutingBgpPeer("address", "key", "value"), `/routing bgp peer set [find remote-address="address"] key="value"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("routing bgp peer mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
