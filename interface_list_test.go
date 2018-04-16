package ros

import (
	"testing"
)

func TestInterfaceList(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{interfaceList("list"), `/interface list print detail where name="list"`},
		{addInterfaceList("list", map[string]string{"comment": "value"}), `:if ([:len [/interface list find name="list"]] = 0) do={/interface list add name="list" comment="value"}`},
		{removeInterfaceList("list"), `/interface list remove [find name="list"]`},
		{setInterfaceList("list", map[string]string{"comment": "value"}), `/interface list set [find name="list"] comment="value"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("interface list mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
