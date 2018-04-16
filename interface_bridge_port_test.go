package ros

import (
	"testing"
)

func TestInterfaceBridgePort(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{interfaceBridgePorts(), `/interface bridge port print detail`},
		{setInterfaceBridgePort("br", "if", "key", "value"), `/interface bridge port set [find bridge="br" interface="if"] key="value"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("interface bridge port mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
