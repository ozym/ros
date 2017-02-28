package ros

import (
	"testing"
)

func TestIPAddress(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{ipAddresses(), `/ip address print detail`},
		{ipAddress("address"), `/ip address print detail where address="address"`},
		{setIPAddress("address", "key", "value"), `/ip address set [find address="address"] key="value"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("ip address mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
