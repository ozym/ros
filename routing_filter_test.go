package ros

import (
	"testing"
)

func TestRoutingFilter(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{routingFilter(map[string]string{"key": "value"}), `/routing filter print detail where key="value"`},
		{addRoutingFilter(map[string]string{"key": "value"}), `:if ([:len [/routing filter find key="value"]] = 0) do={/routing filter add key="value"}`},
		{removeRoutingFilter(map[string]string{"key": "value"}), `/routing filter remove [find key="value"]`},
		{setRoutingFilter(map[string]string{"key": "value"}, map[string]string{"set": "to"}), `/routing filter set [find key="value"] set="to"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("routing filter mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
