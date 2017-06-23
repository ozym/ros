package ros

import (
	"testing"
)

func TestRoutingBGPPeer(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{routingBGPPeer("interface", "address"), `/routing bgp peer print detail where interface="interface" remote-address="address"`},
		{setRoutingBGPPeer("interface", "address", "key", "value"), `/routing bgp peer set [find interface="interface" remote-address="address"] key="value"`},
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
