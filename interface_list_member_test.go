package ros

import (
	"testing"
)

func TestInterfaceListMember(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{interfaceListMember("interface", "list"), `/interface list member print detail where interface="interface" list="list"`},
		{addInterfaceListMember("interface", "list", map[string]string{"comment": "value"}), `:if ([:len [/interface list member find interface="interface" list="list"]] = 0) do={/interface list member add interface="interface" list="list" comment="value"}`},
		{removeInterfaceListMember("interface", "list"), `/interface list member remove [find interface="interface" list="list"]`},
		{setInterfaceListMember("interface", "list", map[string]string{"comment": "value"}), `/interface list member set [find interface="interface" list="list"] comment="value"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("interface list member mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
