package ros

import (
	"testing"
)

func TestInterfaceGre(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{interfaceGre("address"), `/interface gre print detail where remote-address="address"`},
		{setInterfaceGre("address", "key", "value"), `/interface gre set [find remote-address="address"] key="value"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("interface gre mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
